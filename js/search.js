document.addEventListener("DOMContentLoaded", function() {
    const searchInput = document.getElementById("search");
    const artistCards = document.querySelectorAll(".artist-card");

    searchInput.addEventListener("input", function() {
        const searchTerm = searchInput.value.toLowerCase();
        artistCards.forEach(card => {
            const artistName = card.getAttribute("data-name").toLowerCase();
            if (artistName.includes(searchTerm)) {
                card.style.display = "block";
            } else {
                card.style.display = "none";
            }
        });
    });
});