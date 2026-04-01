document.addEventListener("DOMContentLoaded", () => {

    /* ================= 登录 ================= */
    const loginForm = document.querySelector("form[action='/login']");
    if (loginForm) {
        loginForm.addEventListener("submit", async (e) => {
            e.preventDefault();

            const username = loginForm.username.value.trim();
            const password = loginForm.password.value;

            if (!username || !password) {
                return showMessage("请输入用户名和密码", "error");
            }

            try {
                const res = await fetch("/login", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        "Accept": "application/json"
                    },
                    body: JSON.stringify({ username, password })
                });

                const data = await res.json();

                if (res.ok) {
                    showMessage("登录成功，即将跳转...", "success");
                    // replace 防止浏览器回退到登录页
                    window.location.replace("/admin/user");
                } else {
                    showMessage(data.msg || "登录失败", "error");
                }
            } catch (err) {
                console.error(err);
                showMessage("请求异常，请稍后重试", "error");
            }
        });
    }

    /* ================= 注册 ================= */
    const registerForm = document.querySelector("form[action='/register']");
    if (registerForm) {
        registerForm.addEventListener("submit", async (e) => {
            e.preventDefault();

            const username = registerForm.username.value.trim();
            const password = registerForm.password.value;
            const email = registerForm.email.value.trim();
            const ageStr = registerForm.age.value;

            if (!username || !password || !email || !ageStr) {
                return showMessage("请完整填写信息", "error");
            }

            // ⭐ 核心修改：age 字符串 → 数字
            const age = Number(ageStr);

            if (Number.isNaN(age) || age <= 0) {
                return showMessage("年龄必须是有效数字", "error");
            }

            const payload = {
                username,
                password,
                email,
                age   // number
            };

            try {
                const res = await fetch("/register", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        "Accept": "application/json"
                    },
                    body: JSON.stringify(payload)
                });

                const data = await res.json();
                if (res.ok) {
                    showMessage("注册成功，即将跳转到登录页...", "success");
                    setTimeout(() => {
                        // replace 防止浏览器回退到注册页
                        window.location.replace("/login");
                    }, 1000);
                } else {
                    showMessage(data.msg || "注册失败", "error");
                }
            } catch (err) {
                console.error(err);
                showMessage("请求异常，请稍后重试", "error");
            }
        });
    }

    /* ================= 公共提示 ================= */
    function showMessage(msg, type = "error") {
        let container = document.querySelector(".message");
        if (!container) {
            container = document.createElement("div");
            container.className = "message";
            if (loginForm) loginForm.prepend(container);
            if (registerForm) registerForm.prepend(container);
        }
        container.innerText = msg;
        container.className = "message " + type;
    }

    // 防止浏览器后退回登录 / 注册页
    if (window.history.replaceState) {
        window.history.replaceState(null, "", window.location.href);
    }
});
