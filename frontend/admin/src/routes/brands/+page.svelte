<script>
    let searchQuery = $state('');
    let brands = $state([
        { id: 1, name: 'Nike', description: 'Athletic footwear and apparel' },
        { id: 2, name: 'Adidas', description: 'Sports clothing and accessories' },
        { id: 3, name: 'Apple', description: 'Consumer electronics and software' },
        { id: 4, name: 'Samsung', description: 'Electronics, appliances, and mobile devices' },
        { id: 5, name: 'Levi\'s', description: 'Denim jeans and casual wear' }
    ]);
    let filteredBrands = $derived(
        searchQuery ? 
        brands.filter(brand => 
            brand.name.toLowerCase().includes(searchQuery.toLowerCase()) || 
            brand.description.toLowerCase().includes(searchQuery.toLowerCase())
        ) : 
        brands
    );
    
    let showAddModal = $state(false);
    let newBrandName = $state('');
    let newBrandDescription = $state('');
    let addingBrand = $state(false);
    let addError = $state('');
    
    function handleSearch(event) {
        searchQuery = event.target.value;
    }
    
    function openAddModal() {
        showAddModal = true;
        newBrandName = '';
        newBrandDescription = '';
        addError = '';
    }
    
    function closeAddModal() {
        showAddModal = false;
    }
    
    function handleAddBrand(event) {
        event.preventDefault();
        addingBrand = true;
        addError = '';
        
        setTimeout(() => {
            if (!newBrandName.trim()) {
                addError = 'Brand name is required';
                addingBrand = false;
                return;
            }
            
            if (brands.some(brand => brand.name.toLowerCase() === newBrandName.toLowerCase())) {
                addError = 'Brand with this name already exists';
                addingBrand = false;
                return;
            }
            
            const newId = Math.max(...brands.map(b => b.id), 0) + 1;
            brands = [...brands, {
                id: newId,
                name: newBrandName.trim(),
                description: newBrandDescription.trim()
            }];
            
            addingBrand = false;
            showAddModal = false;
        }, 500);
    }
</script>

<div class="flex flex-col min-h-screen p-6" style="background-color: var(--neutral-100);">
    <!-- Header -->
    <div class="w-full mb-6">
        <h1 class="text-2xl font-semibold mb-4" style="color: var(--text-primary);">Brands</h1>
        
        <!-- Search and Add button -->
        <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
            <div class="relative">
                <input
                    type="text"
                    placeholder="Search brands"
                    value={searchQuery}
                    oninput={handleSearch}
                    class="w-full md:w-64 px-3 py-2 border rounded-md"
                    style="border-color: var(--neutral-300); background-color: var(--neutral-50);"
                />
            </div>
            
            <button
                onclick={openAddModal}
                class="px-4 py-2 font-medium rounded-md transition-colors"
                style="background-color: var(--primary-color); color: var(--neutral-800);"
            >
                Add Brand
            </button>
        </div>
    </div>

    <!-- Table of brands -->
    <div class="w-full overflow-x-auto rounded-lg border" style="border-color: var(--neutral-200); background-color: var(--neutral-50);">
        <table class="w-full">
            <thead>
                <tr style="background-color: var(--neutral-200);">
                    <th class="px-4 py-3 text-left" style="color: var(--text-primary);">ID</th>
                    <th class="px-4 py-3 text-left" style="color: var(--text-primary);">Name</th>
                    <th class="px-4 py-3 text-left" style="color: var(--text-primary);">Description</th>
                </tr>
            </thead>
            <tbody>
                {#if filteredBrands.length === 0}
                    <tr>
                        <td colspan="4" class="px-4 py-6 text-center" style="color: var(--text-secondary);">
                            No brands found
                        </td>
                    </tr>
                {/if}
                
                {#each filteredBrands as brand}
                    <tr class="border-t" style="border-color: var(--neutral-200);">
                        <td class="px-4 py-3" style="color: var(--text-secondary);">{brand.id}</td>
                        <td class="px-4 py-3 font-medium" style="color: var(--text-primary);">{brand.name}</td>
                        <td class="px-4 py-3" style="color: var(--text-secondary);">{brand.description}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
    
    <!-- Add Brand Modal -->
    {#if showAddModal}
        <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
            <div class="w-full max-w-md p-6 rounded-lg" style="background-color: var(--neutral-50);">
                <div class="flex justify-between items-center mb-4">
                    <h2 class="text-xl font-semibold" style="color: var(--text-primary);">Add New Brand</h2>
                    <button onclick={closeAddModal} class="text-2xl" style="color: var(--text-secondary);">&times;</button>
                </div>
                
                <form onsubmit={handleAddBrand} class="space-y-4">
                    <div class="space-y-1">
                        <label for=""class="block font-medium" style="color: var(--text-primary);">Brand Name</label>
                        <input
                            type="text"
                            bind:value={newBrandName}
                            class="w-full px-3 py-2 border rounded-md"
                            style="border-color: var(--neutral-300); background-color: var(--neutral-50);"
                            required
                        />
                    </div>
                    
                    <div class="space-y-1">
                        <label for="" class="block font-medium" style="color: var(--text-primary);">Description</label>
                        <textarea
                            bind:value={newBrandDescription}
                            class="w-full px-3 py-2 border rounded-md"
                            style="border-color: var(--neutral-300); background-color: var(--neutral-50);"
                            rows="3"
                        ></textarea>
                    </div>
                    
                    {#if addError}
                        <div class="text-sm" style="color: var(--tertiary-color);">{addError}</div>
                    {/if}
                    
                    <div class="flex justify-end gap-3 pt-2">
                        <button
                            type="button"
                            onclick={closeAddModal}
                            class="px-4 py-2 border rounded-md"
                            style="border-color: var(--neutral-300); color: var(--text-primary);"
                        >
                            Cancel
                        </button>
                        <button
                            type="submit"
                            class="px-4 py-2 font-medium rounded-md transition-colors"
                            style="background-color: var(--primary-color); color: var(--neutral-800);"
                            disabled={addingBrand}
                        >
                            {addingBrand ? 'Adding...' : 'Add Brand'}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    {/if}
</div>