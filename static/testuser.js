

function loadUserList(page = 1, pageSize = 10) {
    fetch(`/admin/users/api?page=${page}&pageSize=${pageSize}`,{
        credentials:"include"
    })
        .then(res => res.json())
        .then(res => {
            console.log("users api:", res)

            const list = res.data || res.list || []
            const total = res.total || res.length

            renderUserTable(list)
            renderPagination(total, page, pageSize)
        })
        .catch(err => {
            console.error(err)
            alert("请求异常")
        })
}


function renderUserTable3(users){
    const tboty = document.getElementById("userTable")
    tboty.innerHTML = ""
    
    users.forEach(user => {
        const tr = document.getElementById("tr")
        tr.innerHTML = `
            <td>${user.ID}</td>
            <td>${user.username}</td>
            <td>${user.age}</td>
            <td>${user.email}</td>
            <td class="actions">
            <button class="btn btn-view" onclick="openUser(${user.ID})">查看</button>
            <button class="btn btn-edit" onclick="openUser(${user.ID})">编辑</button>
            <button class="btn btn-delete" onclick="openUser(${user.ID}, this.closest('tr'))" >删除</button>
            </td>
        `
        tboty.appendChild(tr)
    })
}






































