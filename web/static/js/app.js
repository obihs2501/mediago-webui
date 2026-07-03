// WebSocket connection
let ws = null;
let reconnectInterval = null;

// Connect to WebSocket
function connectWebSocket() {
    const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
    const wsUrl = `${protocol}//${window.location.host}/ws`;

    ws = new WebSocket(wsUrl);

    ws.onopen = () => {
        console.log('WebSocket connected');
        if (reconnectInterval) {
            clearInterval(reconnectInterval);
            reconnectInterval = null;
        }
    };

    ws.onmessage = (event) => {
        const update = JSON.parse(event.data);
        if (update.type === 'task_update') {
            updateTaskInList(update.task);
        }
    };

    ws.onclose = () => {
        console.log('WebSocket disconnected, attempting to reconnect...');
        if (!reconnectInterval) {
            reconnectInterval = setInterval(() => {
                connectWebSocket();
            }, 5000);
        }
    };

    ws.onerror = (error) => {
        console.error('WebSocket error:', error);
    };
}

// Initialize
document.addEventListener('DOMContentLoaded', () => {
    connectWebSocket();
    loadTasks();
    setupEventListeners();
});

function setupEventListeners() {
    // Download form submission
    document.getElementById('downloadForm').addEventListener('submit', async (e) => {
        e.preventDefault();
        await handleDownload();
    });

    // Refresh tasks button
    document.getElementById('refreshBtn').addEventListener('click', () => {
        loadTasks();
    });

    // Load extractors button
    document.getElementById('loadExtractors').addEventListener('click', async () => {
        await loadExtractors();
    });
}

async function handleDownload() {
    const form = document.getElementById('downloadForm');
    const formData = new FormData(form);

    const data = {
        url: formData.get('url'),
        format: formData.get('format') || undefined,
        output: formData.get('output') || undefined,
        cookies: formData.get('cookies') || undefined,
        cookies_browser: formData.get('cookies_browser') || undefined,
        proxy: formData.get('proxy') || undefined,
        yes_playlist: formData.get('yes_playlist') === 'on'
    };

    // Remove undefined fields
    Object.keys(data).forEach(key => {
        if (data[key] === undefined || data[key] === '') {
            delete data[key];
        }
    });

    try {
        const response = await fetch('/api/download', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        if (!response.ok) {
            const error = await response.text();
            throw new Error(error);
        }

        const task = await response.json();
        console.log('Download started:', task);

        // Clear form
        form.reset();

        // Reload tasks
        await loadTasks();

    } catch (error) {
        alert('下载失败: ' + error.message);
        console.error('Download error:', error);
    }
}

async function loadTasks() {
    try {
        const response = await fetch('/api/tasks');
        const tasks = await response.json();

        displayTasks(tasks);
    } catch (error) {
        console.error('Failed to load tasks:', error);
    }
}

function displayTasks(tasks) {
    const tasksList = document.getElementById('tasksList');

    if (!tasks || tasks.length === 0) {
        tasksList.innerHTML = '<p class="empty-state">暂无下载任务</p>';
        return;
    }

    // Sort by created_at descending
    tasks.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));

    tasksList.innerHTML = tasks.map(task => createTaskCard(task)).join('');

    // Attach event listeners
    tasks.forEach(task => {
        const cancelBtn = document.getElementById(`cancel-${task.id}`);
        if (cancelBtn) {
            cancelBtn.addEventListener('click', () => cancelTask(task.id));
        }
    });
}

function createTaskCard(task) {
    const progress = task.progress || 0;
    const statusClass = task.status.toLowerCase();
    const statusText = getStatusText(task.status);

    const createdTime = new Date(task.created_at).toLocaleString('zh-CN');
    const completedTime = task.completed_at
        ? new Date(task.completed_at).toLocaleString('zh-CN')
        : '';

    const showCancel = task.status === 'downloading' || task.status === 'pending';
    const showProgress = task.status === 'downloading';

    return `
        <div class="task-card" data-task-id="${task.id}">
            <div class="task-header">
                <div class="task-title">${task.title || task.url}</div>
                <span class="task-status ${statusClass}">${statusText}</span>
            </div>
            <div class="task-url">${task.url}</div>
            ${showProgress ? `
                <div class="task-progress">
                    <div class="task-progress-bar" style="width: ${progress}%"></div>
                </div>
            ` : ''}
            <div class="task-meta">
                <span>创建: ${createdTime}</span>
                ${completedTime ? `<span>完成: ${completedTime}</span>` : ''}
            </div>
            ${task.error ? `
                <div class="task-error">${escapeHtml(task.error)}</div>
            ` : ''}
            ${showCancel ? `
                <div class="task-actions">
                    <button id="cancel-${task.id}" class="btn-danger">取消</button>
                </div>
            ` : ''}
        </div>
    `;
}

function getStatusText(status) {
    const statusMap = {
        'pending': '等待中',
        'downloading': '下载中',
        'completed': '已完成',
        'failed': '失败',
        'cancelled': '已取消'
    };
    return statusMap[status] || status;
}

function updateTaskInList(task) {
    const taskCard = document.querySelector(`[data-task-id="${task.id}"]`);

    if (taskCard) {
        // Update existing task card
        const tempDiv = document.createElement('div');
        tempDiv.innerHTML = createTaskCard(task);
        taskCard.replaceWith(tempDiv.firstElementChild);

        // Reattach event listener
        const cancelBtn = document.getElementById(`cancel-${task.id}`);
        if (cancelBtn) {
            cancelBtn.addEventListener('click', () => cancelTask(task.id));
        }
    } else {
        // Task doesn't exist, reload all tasks
        loadTasks();
    }
}

async function cancelTask(taskId) {
    if (!confirm('确定要取消此下载任务吗?')) {
        return;
    }

    try {
        const response = await fetch(`/api/tasks/${taskId}/cancel`, {
            method: 'POST'
        });

        if (!response.ok) {
            throw new Error('取消任务失败');
        }

        const task = await response.json();
        updateTaskInList(task);

    } catch (error) {
        alert('取消失败: ' + error.message);
        console.error('Cancel error:', error);
    }
}

async function loadExtractors() {
    const btn = document.getElementById('loadExtractors');
    const container = document.getElementById('extractorsList');

    btn.disabled = true;
    btn.textContent = '加载中...';

    try {
        const response = await fetch('/api/extractors');
        const text = await response.text();

        container.innerHTML = `<pre>${escapeHtml(text)}</pre>`;
    } catch (error) {
        container.innerHTML = `<p class="task-error">加载失败: ${escapeHtml(error.message)}</p>`;
        console.error('Failed to load extractors:', error);
    } finally {
        btn.remove();
    }
}

function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}
