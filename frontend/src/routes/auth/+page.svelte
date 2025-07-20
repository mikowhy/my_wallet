<script lang="ts">
	import { PUBLIC_API_URL } from '$env/static/public';
	let email = '';
	let password = '';
	let message = '';
	let isLoading = false;

	async function handleLogin() {
		isLoading = true;
		message = '';
		try {
			const response = await fetch(`${PUBLIC_API_URL}/auth`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email, password }),
				credentials: 'include'
			});

			const data = await response.json();

			if (response.ok) {
				message = data.message || 'Login successful!';
			} else {
				message = `Error: ${data.message || response.statusText}`;
			}
		} catch (error) {
			message = 'An error occurred. Please try again.';
			console.error('Login error:', error);
		} finally {
			isLoading = false;
		}
	}
</script>

<div class="flex items-center justify-center py-12">
	<div class="w-full max-w-md rounded-lg bg-white p-8 shadow-md">
		<h1 class="mb-6 text-center text-2xl font-bold text-gray-800">Login to My Wallet</h1>
		<form on:submit|preventDefault={handleLogin} class="space-y-6">
			<div>
				<label for="email" class="mb-2 block text-sm font-medium text-gray-700">Email Address</label
				>
				<input
					id="email"
					type="email"
					bind:value={email}
					placeholder="Enter your email"
					class="w-full rounded-md border border-gray-300 p-3 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
					required
				/>
			</div>
			<div>
				<label for="password" class="mb-2 block text-sm font-medium text-gray-700">Password</label>
				<input
					id="password"
					type="password"
					bind:value={password}
					placeholder="Enter your password"
					class="w-full rounded-md border border-gray-300 p-3 shadow-sm focus:border-indigo-500 focus:ring-indigo-500"
					required
				/>
			</div>
			<div>
				<button
					type="submit"
					disabled={isLoading}
					class="w-full rounded-md bg-indigo-600 px-4 py-3 font-semibold text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
				>
					{#if isLoading}Processing...{:else}Login{/if}
				</button>
			</div>
		</form>
		{#if message}
			<p
				class="mt-4 text-center text-sm"
				class:text-red-600={message.startsWith('Error')}
				class:text-green-600={!message.startsWith('Error')}
			>
				{message}
			</p>
		{/if}
	</div>
</div>
