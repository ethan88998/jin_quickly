const params = new URLSearchParams(window.location.search)
const id = params.get("id")
const mode = params.get("mode") || "view" // view / edit

const inputs = ["username", "age", "email"]

fetch(`/admin/user/detail/api?id=${id}`,
    { credentials: "include" })
    .then(res => res.json())
    .then(res => {
        if (res.code !== 200) {
            alert(res.msg)
            return
        }

        const user = res.data
        document.getElementById("username").value = user.username
        document.getElementById("age").value = user.age
        document.getElementById("email").value = user.email

        if (mode === "view") {
            document.getElementById("pageTitle").innerText = "查看用户"
            inputs.forEach(id => document.getElementById(id).disabled = true)
            document.getElementById("saveBtn").style.display = "none"
        } else {
            document.getElementById("pageTitle").innerText = "编辑用户"
        }
    })

function saveUser() {
    fetch(`/admin/user/detail/api?id=${id}`, {
        method: "PUT",
        credentials: "include",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            username: username.value,
            age: Number(age.value),
            email: email.value
        })
    })
        .then(res => res.json())
        .then(res => {
            alert(res.msg)
            if (res.code === 200) goBack()
        })
}

function goBack() {
    window.location.href = "/admin/users"
}
