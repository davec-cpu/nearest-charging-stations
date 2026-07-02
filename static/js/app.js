console.log("hello world!!!");

const map = L.map("map").setView([21.028511, 105.804817], 14);

const chargingStationLayer = L.layerGroup().addTo(map);

L.tileLayer(
    "https://tile.openstreetmap.org/{z}/{x}/{y}.png", 
    {
        maxZoom: 19,
        attribution: "&copy; OpenStreetMap"
    }
).addTo(map);

if (!navigator.geolocation){
    alert("Trình duyệt của bạn không hỗ trợ Geolocation");
} else {
    navigator.geolocation.getCurrentPosition(
        onLocationSuccess,
        onLocationError,
        {
            enableHighAccuracy: true,
            timeout: 5000,
            maximumAge: 0
        }
    )
}

function onLocationSuccess(position){
    const latitude = position.coords.latitude;
    const longitude = position.coords.longitude;

    loadNearbyStations(latitude, longitude);
    console.log("Vị trí hiện tại: ", latitude, longitude);

    map.setView([latitude, longitude], 16);

    L.marker([latitude, longitude])
        .addTo(map)
        .bindPopup("Got a long list of ex-lovers, they're all looking at me like I'm a f***ing psycho")
        .openPopup();
}

function onLocationError(error){
    console.error("Lỗi khi lấy vị trí: ", error.message);
    switch(error.code){
        case error.PERMISSION_DENIED:
            alert("Bạn đã từ chối quyền truy cập vị trí. Vui lòng cho phép để sử dụng tính năng này.");
            break;
        case error.POSITION_UNAVAILABLE:
            alert("Không lấy được vị trí.");
            break;

        case error.TIMEOUT:
            alert("Lấy vị trí quá thời gian.");
            break;

        default:
            alert("Đã xảy ra lỗi.");
    }
}

async function loadNearbyStations(latitude, longitude){
    const radius = 3000;
    const url = `/api/stations/nearby?lat=${latitude}&lng=${longitude}&radius=${radius}`;

    console.log(url);

    try{
        const response = await fetch(url);

        if(!response.ok){
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const stations = await response.json();
        console.log("Các trạm sạc gần đó: ", stations);
        renderChangingStations(stations);

    } catch (error) {
        console.error("Lỗi khi tải các trạm sạc gần đó: ", error);
    }
}
function renderChangingStations(stations){
    chargingStationLayer.clearLayers();
    for (const station of stations){
        const marker = L.marker([station.latitude, station.longitude]);
        marker.bindPopup(`
            <b>${station.name}</b><br>
            <a href="https://www.google.com/maps/search/?api=1&query=${station.latitude},${station.longitude}" target="_blank">Xem trên Google Maps</a>
        `);
        marker.addTo(chargingStationLayer);
    }
}