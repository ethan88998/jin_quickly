document.addEventListener("DOMContentLoaded", function () {
    const params = new URLSearchParams(window.location.search)
    const id = params.get("id")
    const mode = params.get("mode") || "view"

    if (!id) {
        alert("缺少用户 ID")
        return
    }

    const uid = document.getElementById("uid")
    const username = document.getElementById("username")
    const age = document.getElementById("age")
    const email = document.getElementById("email")
    const saveBtn = document.getElementById("saveBtn")
    const pageTitle = document.getElementById("pageTitle")

    // 设置查看/编辑模式
    const isView = mode === "view"
    username.disabled = isView
    age.disabled = isView
    email.disabled = isView
    if (saveBtn) saveBtn.style.display = isView ? "none" : "inline-block"
    pageTitle.innerText = isView ? "查看用户" : "编辑用户"

    // 获取用户数据
    fetch(`/admin/user/detail/api?id=${id}`, { credentials: "include" })
        .then(res => res.json())
        .then(res => {
            if (res.code !== 200) {
                alert(res.msg || "获取失败")
                return
            }
            const u = res.data
            uid.innerText = u.ID
            username.value = u.username
            age.value = u.age
            email.value = u.email
        })
        .catch(err => {
            console.error(err)
            alert("请求异常")
        })

    // 保存用户（编辑模式）
    window.saveUser = function () {
        if (isView) {
            alert("当前为查看模式，不能保存")
            return
        }

        const payload = {
            username: username.value,
            age: Number(age.value),
            email: email.value
        }

        fetch(`/admin/user/detail/api?id=${id}`, {
            method: "PUT",
            credentials: "include",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(payload)
        })
            .then(res => res.json())
            .then(res => {
                alert(res.msg || "保存成功")
                if (res.code === 200) history.back()
            })
            .catch(err => {
                console.error(err)
                alert("保存失败")
            })
    }
})
