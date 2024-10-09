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

// EMAILS

// Form submission handling
const subscribeButtons = document.querySelectorAll('.subscribe-btn, .subscribe-btn-sm');

subscribeButtons.forEach(button => {
    button.addEventListener('click', () => {
        const emailInput = button.closest('.input-group').querySelector('input');
        const email = emailInput.value;

        if (!validateEmail(email)) {
            alert('Խնդրում ենք մուտքագրել վավեր էլ-հասցե։');
            return;
        }

        // Submit the email using Fetch API
        fetch('/email', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email: email }),
        })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert('Շնորհակալություն ձեր բաժանորդագրության համար!');
                emailInput.value = ''; // Clear input
            } else if (data.error === 'This email is already registered.') {
                alert('Այս էլ-հասցեն արդեն գրանցված է։');
                emailInput.value = '';
            } else {
                alert('Բաժանորդագրությունը ձախողվեց, խնդրում ենք փորձել կրկին։');
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Կապի խնդիր, խնդրում ենք փորձել կրկին։');
        });
    });
});

// Email validation function
function validateEmail(email) {
    const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return re.test(String(email).toLowerCase());
}