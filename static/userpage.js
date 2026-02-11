// userpage.js

function renderPagination(total, currentPage, pageSize) {
    const container = document.getElementById("pagination")
    container.innerHTML = ""

    const totalPages = Math.ceil(total / pageSize)
    if (totalPages <= 1) return

    for (let i = 1; i <= totalPages; i++) {
        const btn = document.createElement("button")
        btn.innerText = i

        if (i === currentPage) {
            btn.disabled = true
        }

        btn.onclick = () => loadUserList(i, pageSize)
        container.appendChild(btn)
    }
}
