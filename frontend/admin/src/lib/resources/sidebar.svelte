<script>
    import { House, LayoutGrid, Tag, User, Box, ChartColumnDecreasing, ChartSpline, FileChartColumn, Sun, Library, Tags, Radical } from "lucide-svelte";

    const MenuItems = [
        { name: "Home", url: "/", icon: House },
        { name: "Sections", url: "/sections", icon: Library },
        { name: "Categories", url: "/categories", icon: LayoutGrid },
        { name: "Brands", url: "/brands", icon: Tags },
        { name: "Products", url: "/products", icon: Box },
        { name: "Customers", url: "/customers", icon: User },
        { name: "Stocks", url: "/stocks", icon: ChartSpline },
        { name: "Sales", url: "/sales", icon: FileChartColumn },
        { name: "Analytics", url: "/analytics", icon: Radical},
    ];

    let activeItem = "Home";

    function setActive(name) {
        activeItem = name;
    }
</script>

<style>
    .sidebar {
        width: 60px;
        height: 100vh;
        background-color: var(--primary-color);
        display: flex;
        flex-direction: column;
        align-items: center;
        padding: 16px 0;
        position: fixed;
        left: 0;
        top: 0;
        z-index: 100;
        transition: width 0.3s;
    }

    .menu-items {
        display: flex;
        flex-direction: column;
        width: 100%;
        gap: 8px;
    }

    .menu-item {
        width: 100%;
        display: flex;
        align-items: center;
        padding: 10px 0;
        color: var(--soft-gray);
        position: relative;
        cursor: pointer;
        transition: all 0.2s;
    }

    .menu-item:hover {
        color: var(--secondary-color);
    }

    .menu-item.active {
        color: var(--secondary-color);
    }

    .menu-item.active::before {
        content: "";
        position: absolute;
        left: 0;
        top: 0;
        width: 3px;
        height: 100%;
        background-color: var(--tertiary-color);
    }

    .icon-container {
        width: 60px;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .tooltip {
        position: absolute;
        left: 60px;
        background-color: var(--secondary-color);
        color: white;
        padding: 6px 12px;
        border-radius: 4px;
        white-space: nowrap;
        opacity: 0;
        visibility: hidden;
        transition: opacity 0.2s, visibility 0.2s;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
        z-index: 101;
    }

    .menu-item:hover .tooltip {
        opacity: 1;
        visibility: visible;
    }

    /* Add a small triangle to the tooltip */
    .tooltip::before {
        content: "";
        position: absolute;
        top: 50%;
        left: -5px;
        transform: translateY(-50%);
        border-width: 5px 5px 5px 0;
        border-style: solid;
        border-color: transparent var(--secondary-color) transparent transparent;
    }
</style>

<div class="sidebar">    
    <div class="menu-items">
        {#each MenuItems as item}
            <button 
                class="menu-item {activeItem === item.name ? 'active' : ''}" 
                on:click={() => setActive(item.name)}
                on:keydown={(e) => e.key === 'Enter' && setActive(item.name)}
                role="menuitem"
            >
                <div class="icon-container">
                    <svelte:component this={item.icon} size={20} />
                </div>
                <div class="tooltip">{item.name}</div>
            </button>
        {/each}
    </div>
</div>