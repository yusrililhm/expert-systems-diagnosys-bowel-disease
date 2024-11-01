function changePassword() {
    const data = new FormData(document.getElementById("change_password"));

    fetch('/api/v1/users/change-password', {
        method: 'PATCH',
        body: data,
    })
    .then(() => {
        if (confirm('Password berhasil diubah')) {
         window.location.href = "/profile"   
        }
    })
    .catch(e => {
        console.error(e)
    })
}
