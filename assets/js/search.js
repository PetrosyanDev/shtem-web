// // Get all sections with class "col-12"
// const sections = document.querySelectorAll('.col-12');
// let shtems = []

// // Loop through each section
// sections.forEach(section => {
//     const itemsContainer = section.querySelector('.items');

//     if (itemsContainer) {
//         const items = itemsContainer.querySelectorAll('a');
//         const category = section.querySelector('h2').textContent;

//         items.forEach(item => {

//             const href = item.getAttribute('href');
//             const name = item.querySelector('.shtem-name').textContent;

//             shtems.push({name: name, link: href, category: category, element: item})
//         });
//     }
// });

// const searchInput = document.querySelector('[data-search]')

// searchInput.addEventListener("input", e => {
//     const value = e.target.value.toLowerCase();;
//     shtems.forEach(shtem => {
//         const isVisible = shtem.name.toLowerCase().includes(value) || shtem.link.toLowerCase().includes(value) || shtem.category.toLowerCase().includes(value)
//         shtem.element.classList.toggle("d-none", !isVisible)
//     })
// })
// Get all sections with class "col-12"
const sections = document.querySelectorAll('.col-12');
let shtems = [];

// Loop through each section
sections.forEach(section => {
    const itemsContainer = section.querySelector('.items');

    if (itemsContainer) {
        const items = itemsContainer.querySelectorAll('a');
        const category = section.querySelector('h2').textContent;

        items.forEach(item => {
            const href = item.getAttribute('href');
            const name = item.querySelector('.shtem-name').textContent;

            shtems.push({ name: name, link: href, category: category, element: item });
        });

    }
});

const searchInput = document.querySelector('[data-search]');

searchInput.addEventListener("input", e => {
    const value = e.target.value.toLowerCase();
    shtems.forEach(shtem => {
        const isVisible = shtem.name.toLowerCase().includes(value) || shtem.link.toLowerCase().includes(value) || shtem.category.toLowerCase().includes(value);
        shtem.element.classList.toggle("d-none", !isVisible);
    });

    // Update visibility of sections based on the visibility of their items
    sections.forEach(section => {
        const items = section.querySelectorAll('.items a');
        let isVisible = false;

        items.forEach(item => {
            if (!item.classList.contains('d-none')) {
                isVisible = true;
            }
        });

        section.classList.toggle("d-none", !isVisible);
    });
});
