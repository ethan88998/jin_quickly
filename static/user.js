// user.js

function loadUserList(page = 1, pageSize = 10) {
    fetch(`/admin/user/api?page=${page}&pageSize=${pageSize}`, {
        credentials: "include"
    })
        .then(res => res.json())
        .then(res => {
            console.log("users api:", res)

            const list = res.data || res.list || []
            const total = res.total || list.length

            renderUserTable(list)
            renderPagination(total, page, pageSize)
        })
        .catch(err => {
            console.error(err)
            alert("请求异常")
            alert("请求异常")
        })
}


function renderUserTable(users) {
    const tbody = document.getElementById("userTable")
    tbody.innerHTML = ""

    users.forEach(user => {
        const tr = document.createElement("tr")
        tr.innerHTML = `
            <td>${user.ID}</td>
            <td>${user.username}</td>
            <td>${user.age}</td>
            <td>${user.email}</td>
            <td>${formatTime(user.CreatedAt)}</td>
            <td>${formatTime(user.UpdatedAt)}</td>
            <td class="actions">
                <button class="btn btn-view"  onclick="openUser(${user.ID}, 'view')">查看</button>
                <button class="btn btn-edit"  onclick="openUser(${user.ID}, 'edit')">编辑</button>
                <button class="btn btn-delete"  onclick="deleteUser(${user.ID}, this.closest('tr'))">删除</button>
            </td>
        `
        tbody.appendChild(tr)
    })
}

function openUser(id, mode) {
    location.href = `/admin/user/api/detail?id=${id}&mode=${mode}`
}

function deleteUser(id, row) {
    if (!confirm("确认删除？")) return

    fetch(`/admin/user/${id}`, { method: "DELETE", credentials: "include" })
        .then(res => res.json())
        .then(res => {
            if (res.code === 200) {
                row.remove()
            }
            alert(res.msg)
        })
}

// 用户统计
function loadUserStat() {
fetch('/admin/user/api/total')
    .then(res => res.json())
    .then(data => {
        document.getElementById('totalUsers').innerText = data.total
        document.getElementById('todayUsers').innerText = data.today
    })
    .catch(err => {
        console.error('统计加载失败', err)
    })
}

// 注册时间/更新时间格式处理
function formatTime(timeStr) {
    const date = new Date(timeStr);
    const Y = date.getFullYear();
    const M = String(date.getMonth() + 1).padStart(2, '0');
    const D = String(date.getDate()).padStart(2, '0');
    const h = String(date.getHours()).padStart(2, '0');
    const m = String(date.getMinutes()).padStart(2, '0');
    const s = String(date.getSeconds()).padStart(2, '0');
    return `${Y}-${M}-${D} ${h}:${m}:${s}`;
}


// 全局搜索/分页状态
const currentParams = {
    page: 1,
    pageSize: 10,
    username: '',
    age: '',
    start_date: '',
    end_date: ''
}



// 唯一请求函数：loadUsers
function loadUsers() {
    const params = new URLSearchParams()

    Object.keys(currentParams).forEach(key => {
        if (currentParams[key] !== '' && currentParams[key] !== null) {
            params.append(key, currentParams[key])
        }
    })

    fetch('/admin/user/api?' + params.toString())
        .then(res => res.json())
        .then(res => {
            renderUserTable(res.list)   // ✅ 修正这里
            renderPagination(res.page, res.total)
        })
        .catch(err => {
            console.error('加载用户失败', err)
        })
}

// 搜索
function searchUsers() {
    currentParams.page = 1
    currentParams.username = document.getElementById('searchUsername').value.trim()
    currentParams.age = document.getElementById('searchAge').value.trim()
    currentParams.start_date = document.getElementById('startDate').value
    currentParams.end_date = document.getElementById('endDate').value

    loadUsers()
}

// 分页
function renderPagination(page, total) {
    const container = document.getElementById('pagination')
    container.innerHTML = ''

    const totalPages = Math.ceil(total / currentParams.pageSize)

    for (let i = 1; i <= totalPages; i++) {
        const btn = document.createElement('button')
        btn.innerText = i
        if (i === page) btn.disabled = true

        btn.onclick = () => {
            currentParams.page = i
            loadUsers()
        }

        container.appendChild(btn)
    }
}

// 参数清空
function resetSearch() {
    currentParams.page = 1
    currentParams.username = ''
    currentParams.age = ''
    currentParams.start_date = ''
    currentParams.end_date = ''

    document.getElementById('searchUsername').value = ''
    document.getElementById('searchAge').value = ''
    document.getElementById('startDate').value = ''
    document.getElementById('endDate').value = ''

    loadUsers()
}

document.addEventListener('DOMContentLoaded', () => {
    loadUsers()
})
