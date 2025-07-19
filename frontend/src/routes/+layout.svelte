<script>
	import '../app.css';
</script>

<div class="min-h-screen bg-gray-50">
	<header class="bg-white shadow-sm">
		<nav class="container mx-auto flex items-center justify-between p-4">
			<a href="/" class="text-xl font-bold text-indigo-600">My Wallet</a>
			<div class="space-x-6">
				<a href="/" class="text-sm font-medium text-gray-500 hover:text-gray-900">Home</a>
				<a href="/auth" class="text-sm font-medium text-gray-500 hover:text-gray-900">Login</a>
			</div>
		</nav>
	</header>

	<main>
		<slot />
	</main>
</div>