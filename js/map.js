        // Initialiser la carte
        var map = L.map('map').setView([51.505, -0.09], 2);

        // Ajouter les tuiles OpenStreetMap
        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        }).addTo(map);

        // Ajouter les marqueurs pour les locations
        var locations = JSON.parse('{{.LocationsJSON}}');
        if (locations && locations.length === 0) {
            document.querySelector('ul').innerHTML = '<li>Aucune location disponible</li>';
        } else if (locations) {
            locations.forEach(function(location) {
                var marker = L.marker([location.Lat, location.Lng]).addTo(map);
                marker.bindPopup(location.Name).openPopup();
            });
        }