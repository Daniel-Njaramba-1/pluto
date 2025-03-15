<script lang="js">
    let email = $state('');
    let username = $state('');
    let password = $state('');
    let confirmPassword = $state('');
    let registerLoading = $state(false);
    let registerError = $state('');

    async function handleRegister(event) {
        event.preventDefault();
        registerLoading = true;
        registerError = '';

        if (!email.includes('@')) {
            registerError = 'Please enter a valid email address';
            registerLoading = false;
            return;
        }

        if (password.length < 6) {
            registerError = 'Password must be at least 6 characters';
            registerLoading = false;
            return;
        }

        if (password !== confirmPassword) {
            registerError = 'Passwords do not match';
            registerLoading = false;
            return;
        }

        console.log('Logging in with:', { username, email, password });

        try {
            let response = await fetch('http://localhost:8080/api/admin/register', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    email,
                    username,
                    password,
                }),
            });

            let result = await response.json();
            if (response.ok) {
                console.log('Registration successful:', result);
            } else {
                registerError = result.error || 'Registration failed';
            }
        } catch (error) {
            registerError = 'Network error';
        } finally {
            registerLoading = false;
        }
        
    }
</script>

<div class="flex items-center justify-center min-h-screen" style="background-color: var(--neutral-100);">
    <div class="w-full max-w-[400px] p-6 border rounded-lg shadow-md" style="background-color: var(--neutral-50); border-color: var(--neutral-200);">
        <div class="mb-6">
            <h1 class="text-2xl" style="color: var(--text-primary);">Register</h1>
        </div>

        <form onsubmit={handleRegister} class="space-y-6">
            <div class="space-y-1">
                <label for="email" class="block font-medium" style="color: var(--text-primary);">
                    Email
                </label>
                
                <input
                    id="email"
                    type="email"
                    bind:value={email}
                    placeholder=""
                    class="w-full px-3 py-2 border rounded-md" 
                    style="border-color: var(--neutral-300); background-color: var(--neutral-50);"
                    required
                />
            </div>

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

            <div>
                <label for="password" class="block font-medium" style="color: var(--text-primary);">
                    Password
                </label>

                <input
                    id="password"
                    type="password"
                    bind:value={password}
                    placeholder=""
                    class="w-full px-3 py-2 border rounded-md"
                    style="border-color: var(--neutral-300); background-color: var(--neutral-50);"
                    required
                />
            </div>

            <div>
                <label for="confirmPassword" class="block font-medium" style="color: var(--text-primary);">
                    Confirm Password
                </label>
                <input
                    id="confirmPassword"
                    type="password"
                    bind:value={confirmPassword}
                    placeholder=""
                    class="w-full px-3 py-2 border rounded-md"
                    style="border-color: var(--neutral-300); background-color: var(--neutral-50);"
                    required
                />
            </div>

            <!-- Error Message -->
            {#if registerError}
                <div class="text-sm text-red-500">{registerError}</div>
            {/if}

            <!-- Submit Button -->
            <div class="flex space-x-4">
                <button
                    type="submit"
                    class="py-2 px-4 font-medium rounded-md transition-colors"
                    style="background-color: var(--primary-color); color: var(--neutral-800);"
                    disabled={registerLoading}
                >
                    {registerLoading ? 'Registering...' : 'Register'}
                </button>
            </div>


            <!-- Login Link -->
            <p class="text-center text-sm text-gray-600">
                Already have an account?
                <a href="/auth/login" class="font-medium text-blue-600 hover:underline">
                    Sign in
                </a>
            </p>
        </form>
    </div>
</div>