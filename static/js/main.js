// ======== 菜单切换 ========
document.querySelectorAll('.sidebar li').forEach(li => {
    li.addEventListener('click', function() {
        document.querySelectorAll('.sidebar li').forEach(item => item.classList.remove('active'))
        li.classList.add('active')
        const page = li.dataset.page
        loadPage(page)
    })
})

// ======== 加载内容 ========
function loadPage(page) {
    const content = document.getElementById('mainContent')
    if(page === 'users'){
        content.innerHTML = `
            <h2>用户管理</h2>
            <table>
                <thead>
                    <tr>
                        <th>ID</th><th>用户名</th><th>年龄</th><th>邮箱</th><th>操作</th>
                    </tr>
                </thead>
                <tbody id="userTable"></tbody>
            </table>
            <button onclick="showAddUserForm()">新增用户</button>
        `
        loadUserList()
    } else {
        content.innerHTML = `<h2>${page} 页面开发中...</h2>`
    }
}

// ======== 用户列表 ========
function loadUserList(){
    fetch("/admin/users/api", {credentials: "include"})
        .then(res => res.json())
        .then(res => {
            const tbody = document.getElementById("userTable")
            tbody.innerHTML = ""
            res.data.forEach(user => {
                const tr = document.createElement("tr")
                tr.innerHTML = `
                    <td>${user.ID}</td>
                    <td>${user.username}</td>
                    <td>${user.age}</td>
                    <td>${user.email}</td>
                    <td>
                        <button class="btn-view" onclick="viewUser(${user.ID})">查看</button>
                        <button class="btn-edit" onclick="editUser(${user.ID})">编辑</button>
                        <button class="btn-delete" onclick="deleteUser(${user.ID}, this.closest('tr'))">删除</button>
                    </td>
                `
                tbody.appendChild(tr)
            })
        })
}

// ======== 编辑 / 查看 / 删除 / 新增 ========
function editUser(id) { window.location.href = `/admin/user/edit?id=${id}` }
function viewUser(id) { alert("查看用户：" + id) }
function deleteUser(id, row) { alert("删除用户：" + id) }

function showAddUserForm() {
    const content = document.getElementById('mainContent')
    content.innerHTML = `
        <h2>新增用户</h2>
        <div class="card">
            <div class="form-row">
                <label>用户名：</label><input id="username">
            </div>
            <div class="form-row">
                <label>年龄：</label><input id="age" type="number">
            </div>
            <div class="form-row">
                <label>邮箱：</label><input id="email">
            </div>
            <div class="actions">
                <button class="btn-cancel" onclick="loadPage('users')">取消</button>
                <button class="btn-save" onclick="submitAddUser()">保存</button>
            </div>
        </div>
    `
}
