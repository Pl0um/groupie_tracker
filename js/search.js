// Attendre que le DOM soit complètement chargé avant d'exécuter le script
document.addEventListener("DOMContentLoaded", function() {
    // Récupérer l'élément de saisie de recherche et les cartes d'artistes
    const searchInput = document.getElementById("search");
    const artistCards = document.querySelectorAll(".artist-card");

    // Ajouter un écouteur d'événement pour détecter les changements dans le champ de recherche
    searchInput.addEventListener("input", function() {
        const searchTerm = searchInput.value.toLowerCase();
        artistCards.forEach(card => {
            const artistName = card.getAttribute("data-name").toLowerCase();
            // Afficher ou masquer les cartes en fonction du terme de recherche
            if (artistName.includes(searchTerm)) {
                card.style.display = "block";
            } else {
                card.style.display = "none";
            }
        });
    });
});

// Initialiser les ordres de tri
let sortOrder = "asc";
let sortAlphaOrder = "asc";

// Fonction pour basculer l'ordre de tri par date de création
function toggleSortOrder() {
    sortOrder = sortOrder === "asc" ? "desc" : "asc";
    sortArtists(sortOrder);
    updateSortArrow();
}

// Fonction pour trier les artistes par date de création
function sortArtists(order) {
    const artistsGrid = document.getElementById("artists-grid");
    const artistCards = Array.from(artistsGrid.getElementsByClassName("artist-card"));

    artistCards.sort((a, b) => {
        // Comparer les dates de création des artistes
        const dateA = parseInt(a.getAttribute("data-creation-date"));
        const dateB = parseInt(b.getAttribute("data-creation-date"));
        return order === "asc" ? dateA - dateB : dateB - dateA;
    });

    // Réorganiser les cartes d'artistes dans la grille
    artistCards.forEach(card => artistsGrid.appendChild(card));
}

// Fonction pour basculer l'ordre de tri alphabétique
function toggleSortAlphaOrder() {
    sortAlphaOrder = sortAlphaOrder === "asc" ? "desc" : "asc";
    sortArtistsAlpha(sortAlphaOrder);
    updateSortAlphaArrow();
}

// Fonction pour trier les artistes par ordre alphabétique
function sortArtistsAlpha(order) {
    const artistsGrid = document.getElementById("artists-grid");
    const artistCards = Array.from(artistsGrid.getElementsByClassName("artist-card"));

    artistCards.sort((a, b) => {
        // Comparer les noms des artistes
        const nameA = a.getAttribute("data-name").toLowerCase();
        const nameB = b.getAttribute("data-name").toLowerCase();
        if (order === "asc") {
            return nameA.localeCompare(nameB);
        } else {
            return nameB.localeCompare(nameA);
        }
    });

    // Réorganiser les cartes d'artistes dans la grille
    artistCards.forEach(card => artistsGrid.appendChild(card));
}

// Fonction pour mettre à jour la flèche de tri par date de création
function updateSortArrow() {
    const sortArrow = document.getElementById("sort-arrow");
    sortArrow.textContent = sortOrder === "asc" ? "⬇️" : "⬆️";
}

// Fonction pour mettre à jour la flèche de tri alphabétique
function updateSortAlphaArrow() {
    const sortAlphaArrow = document.getElementById("sort-alpha-arrow");
    sortAlphaArrow.textContent = sortAlphaOrder === "asc" ? "⬇️" : "⬆️";
}