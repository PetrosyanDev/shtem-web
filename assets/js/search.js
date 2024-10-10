document.addEventListener("DOMContentLoaded", () => {
    // Search input element
    const searchInput = document.querySelector('[data-search]');
    
    // Prevent form submission and handle the "Enter" key
    searchInput.addEventListener('keydown', function (e) {
        if (e.key === 'Enter') {
            e.preventDefault(); // Prevent default form submission
            const searchTerm = searchInput.value.trim(); // Get search input and trim extra spaces

            if (searchTerm) {
                window.location.href = `/?q=${encodeURIComponent(searchTerm)}`; // Redirect with query parameter
            }
        }
    });

    // Get all sections with class "col-12"
    const sections = document.querySelectorAll('.col-12');
    let shtems = [];
    const subed = document.querySelector('.subscribe-section');
    const noResultsMessage = document.querySelector('#no-found');

    // Set "no results" message and hide it initially
    noResultsMessage.textContent = "Շտեմարաններ չեն գտնվել";
    noResultsMessage.classList.add('no-results', 'd-none');

    // Extract search query from URL
    const urlParams = new URLSearchParams(window.location.search);
    const query = urlParams.get('q') ? urlParams.get('q').toLowerCase() : '';

    // Pre-fill search input with the query if available
    if (query) {
        searchInput.value = query;
    }

    // Check if sections exist
    if (sections.length > 0) {
        // Loop through each section and gather items for filtering
        sections.forEach(section => {
            const itemsContainer = section.querySelector('.items');

            if (itemsContainer) {
                const items = itemsContainer.querySelectorAll('a');
                const category = section.querySelector('h2').textContent;

                items.forEach(item => {
                    const href = item.getAttribute('href');
                    const name = item.querySelector('.shtem-name').textContent;

                    // Add items to the shtems array for filtering
                    shtems.push({ name: name, link: href, category: category, element: item });
                });
            }
        });

        // Apply filtering if there's a search query in the URL
        if (query) {
            filterShtems(query);
        }

        // Listen for input changes to filter results
        searchInput.addEventListener("input", e => {
            const value = e.target.value.toLowerCase();
            filterShtems(value);
        });
    }

    // Function to filter the shtems based on search value
    function filterShtems(value) {
        let anyVisible = false;

        // Loop through the shtems array and check visibility
        shtems.forEach(shtem => {
            const isVisible = shtem.name.toLowerCase().includes(value) || shtem.link.toLowerCase().includes(value) || shtem.category.toLowerCase().includes(value);
            shtem.element.classList.toggle("d-none", !isVisible);
            if (isVisible) {
                anyVisible = true;
            }
        });

        // Show or hide the "no results" message and subscribe section
        noResultsMessage.classList.toggle('d-none', anyVisible);
        subed.classList.toggle('d-none', !anyVisible);

        // Update section visibility based on the visibility of their items
        sections.forEach(section => {
            const items = section.querySelectorAll('.items a');
            let sectionVisible = false;

            items.forEach(item => {
                if (!item.classList.contains('d-none')) {
                    sectionVisible = true;
                }
            });

            section.classList.toggle("d-none", !sectionVisible);
        });
    }

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
});