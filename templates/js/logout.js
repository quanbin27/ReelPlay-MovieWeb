function logout() {
    // Xóa token khỏi localStorage
    localStorage.removeItem("token");
    // Điều hướng về trang đăng nhập hoặc trang chính
    window.location.href = "/signin";
}