function submitUser() {
    const data = {
        username: username.value.trim(),
        password: password.value.trim(),
        age: Number(age.value),
        email: email.value.trim()
    }

    if (!data.username || !data.password) {
        alert("用户名和密码必填")
        return
    }

    fetch('/admin/user/add/api', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify(data)
    })
        .then(res => res.json())
        .then(res => {
            alert(res.msg)
            if (res.code === 200) {
                location.href = "/admin/user"
            }
        })
}
