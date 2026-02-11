


function searchUsers(page = 1) {
    const username = document.getElementById("searchUsername").value.trim()
    const age = document.getElementById("searchAge").value.trim()
    const startDate = document.getElementById("startDate").value
    const  endDate = document.getElementById("endDate").value

    const params = new URLSearchParams({
        page: page,
        pageSize: 10
    })

    if (username) params.append("username", username)
    if (age) params.append("age", age)
    if (startDate) params.append("start_date", startDate)
    if (endDate) params.append("end_date", endDate)

    fetch(`/admin/user/search/api?`+ params.toString())
        .then(res => res.json())
        .then(res => {
            renderTable9(res.data)
            renderPagination(1, res.total)
        })
        .catch(err => {
            console.error("搜索失败", err)
        })
}

// 清除搜索注册日期输入框
function resetSearch() {
    document.getElementById('searchUsername').value = ''
    document.getElementById('searchAge').value = ''
    document.getElementById('startDate').value = ''
    document.getElementById('endDate').value = ''
    loadUserList(1, 10)
}
