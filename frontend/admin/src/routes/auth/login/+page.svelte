<script>
    let username = $state('');
    let password = $state('');
    let loginLoading = $state(false);
    let loginError = $state('');

    function handleLogin(event) {
        event.preventDefault();
        loginLoading = true;
        loginError = '';

        setTimeout(() => {
            if (password.length < 6) {
                loginError = 'Password must be at least 6 characters';
                loginLoading = false;
                return;
            }

            console.log('Logging in with:', { username, password });
            loginLoading = false;
        }, 1000);
    }
</script>

<div class="flex items-center justify-center min-h-screen" style="background-color: var(--neutral-100);">
    <div class="w-full max-w-[400px] p-6 border rounded-lg shadow-md" style="background-color: var(--neutral-50); border-color: var(--neutral-200);">
        <div class="mb-6">
            <h1 class="text-2xl" style="color: var(--text-primary);">Login</h1>
        </div>
    
        <form onsubmit={handleLogin} class="space-y-4">
            <div class="space-y-1">
                <label for="username" class="block font-medium" style="color: var(--text-primary);">
                    Username
                </label>
                
                <input
                    id="username"
                    type="text"
                    bind:value={username}
                    placeholder=""
                    class="w-full px-3 py-2 border rounded-md" 
                    style="border-color: var(--neutral-300); background-color: var(--neutral-50);"
                    required
                />
            </div>
        
            <div class="space-y-1">
                <div class="flex justify-between items-center">
                    <label for="password" class="block font-medium" style="color: var(--text-primary);">
                        Password
                    </label>
                    <a href="#top" class="text-sm hover:underline text-blue-600">
                        Forgot your password?
                    </a>
                </div>
                
                <input
                    id="password"
                    type="password"
                    bind:value={password}
                    class="w-full px-3 py-2 border rounded-md"
                    style="border-color: var(--neutral-300); background-color: var(--neutral-50);"
                    required
                />
            </div>
        
            {#if loginError}
                <div class="text-sm" style="color: var(--tertiary-color);">{loginError}</div>
            {/if}
        
            <div class="flex space-x-4">
                <button
                    type="submit"
                    class="py-2 px-4 font-medium rounded-md transition-colors"
                    style="background-color: var(--primary-color); color: var(--neutral-800);"
                    disabled={loginLoading}
                >
                    {loginLoading ? 'Logging in...' : 'Login'}
                </button>
            </div>

        
            <div class="mt-6 text-center">
                <span style="color: var(--text-secondary);">Don't have an account? </span>
                <a href="#top" class="font-medium hover:underline text-blue-600">
                    Sign up
                </a>
            </div>
        </form>
    </div>
</div>