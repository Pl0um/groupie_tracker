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

let sortOrder = "asc";
let sortAlphaOrder = "asc";

function toggleSortOrder() {
    sortOrder = sortOrder === "asc" ? "desc" : "asc";
    sortArtists(sortOrder);
    updateSortArrow();
}

function sortArtists(order) {
    const artistsGrid = document.getElementById("artists-grid");
    const artistCards = Array.from(artistsGrid.getElementsByClassName("artist-card"));

    artistCards.sort((a, b) => {
        const dateA = parseInt(a.getAttribute("data-creation-date"));
        const dateB = parseInt(b.getAttribute("data-creation-date"));

        if (order === "asc") {
            return dateA - dateB;
        } else {
            return dateB - dateA;
        }
    });

    artistCards.forEach(card => {
        artistsGrid.appendChild(card);
    });
}

function toggleSortAlphaOrder() {
    sortAlphaOrder = sortAlphaOrder === "asc" ? "desc" : "asc";
    sortArtistsAlpha(sortAlphaOrder);
    updateSortAlphaArrow();
}

function sortArtistsAlpha(order) {
    const artistsGrid = document.getElementById("artists-grid");
    const artistCards = Array.from(artistsGrid.getElementsByClassName("artist-card"));

    artistCards.sort((a, b) => {
        const nameA = a.getAttribute("data-name").toLowerCase();
        const nameB = b.getAttribute("data-name").toLowerCase();

        if (order === "asc") {
            return nameA.localeCompare(nameB);
        } else {
            return nameB.localeCompare(nameA);
        }
    });

    artistCards.forEach(card => {
        artistsGrid.appendChild(card);
    });
}

function updateSortArrow() {
    const sortArrow = document.getElementById("sort-arrow");
    if (sortOrder === "asc") {
        sortArrow.textContent = "⬇️";
    } else {
        sortArrow.textContent = "⬆️";
    }
}

function updateSortAlphaArrow() {
    const sortAlphaArrow = document.getElementById("sort-alpha-arrow");
    if (sortAlphaOrder === "asc") {
        sortAlphaArrow.textContent = "⬇️";
    } else {
        sortAlphaArrow.textContent = "⬆️";
    }
}

function toggleFilters() {
    const filterOptions = document.getElementById("filter-options");
    if (filterOptions.style.display === "none") {
        filterOptions.style.display = "block";
    } else {
        filterOptions.style.display = "none";
    }
}