const apiUrl = 'http://localhost:8080';

document.addEventListener('DOMContentLoaded', () => {
    const signupForm = document.getElementById('signupForm');
    const loginForm = document.getElementById('loginForm');
    const message = document.getElementById('message');

    if (signupForm) {
        signupForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const res = await fetch(`${apiUrl}/signup`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    email: document.getElementById('signupEmail').value,
                    password: document.getElementById('signupPassword').value,
                    role: document.getElementById('signupRole').value
                })
            });
            const data = await res.json();
            message.innerText = JSON.stringify(data);
        });
    }

    if (loginForm) {
        loginForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const res = await fetch(`${apiUrl}/login`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    email: document.getElementById('loginEmail').value,
                    password: document.getElementById('loginPassword').value
                })
            });
            const data = await res.json();
            if (data.token) {
                localStorage.setItem('jwt', data.token);
                window.location.href = '/profile';
            } else {
                message.innerText = JSON.stringify(data);
            }
        });
    }

    const profileDiv = document.getElementById('profile');
    const logoutBtn = document.getElementById('logoutBtn');

    if (profileDiv) {
        const token = localStorage.getItem('jwt');
        if (!token) { window.location.href = 'index.html'; }
        fetch(`${apiUrl}/profile`, { headers: { 'Authorization': `Bearer ${token}` }})
            .then(res => res.json())
            .then(data => profileDiv.innerText = JSON.stringify(data, null, 2))
            .catch(() => window.location.href = 'index.html');
    }

    if (logoutBtn) {
        logoutBtn.addEventListener('click', () => {
            localStorage.removeItem('jwt');
            window.location.href = '/';
        });
    }
});
