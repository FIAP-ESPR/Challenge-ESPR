const mapa = L.map("map").setView([-23.5505, -46.6333], 13); // SP como centro

L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
  attribution: "&copy; OpenStreetMap contributors",
}).addTo(mapa);

// Coordenadas fixas para pontos fictícios
const pontos = {
    "Armazém A": [-23.5505, -46.6333],
    "Centro B": [-23.5530, -46.6200],
    "Depósito C": [-23.5400, -46.6300],
    "Loja D": [-23.5450, -46.6400],
    "Posto E": [-23.5300, -46.6250],
    "CD F": [-23.5550, -46.6450],
    "Estoque G": [-23.5280, -46.6350],
    "Distribuidor H": [-23.5480, -46.6500],
    "Retirada I": [-23.5250, -46.6400],
    "Hub K": [-23.5600, -46.6200],
    "Filial J": [-23.5380, -46.6480],
    "Depósito L": [-23.5330, -46.6600],
    "Centro M": [-23.5470, -46.6150],
    "Galpão N": [-23.5340, -46.6650],
    "Unidade O": [-23.5555, -46.6100],
    "Remessa P": [-23.5305, -46.6700],
    "CD Q": [-23.5570, -46.6050],
    "Armazém R": [-23.5260, -46.6750],
    "Base S": [-23.5590, -46.6000],
    "Estoque T": [-23.5240, -46.6800],
    "Logística U": [-23.5605, -46.5950],
    "Loja V": [-23.5225, -46.6850],
    "Centro W": [-23.5615, -46.5900],
    "Filial X": [-23.5210, -46.6900],
    "Depósito Y": [-23.5625, -46.5850],
    "Galpão Z": [-23.5195, -46.6950],
    "Posto AA": [-23.5635, -46.5800],
    "Unidade AB": [-23.5180, -46.7000],
    "CD AC": [-23.5645, -46.5750],
    "Ponto AD": [-23.5165, -46.7050],
    "Hub AE": [-23.5655, -46.5700],
  };
  

Object.entries(pontos).forEach(([nome, coords]) => {
  L.marker(coords).addTo(mapa).bindPopup(nome);
});

let rotaLayer = null;

async function calcularRota() {
  const from = document.getElementById("from").value;
  const to = document.getElementById("to").value;

  const res = await fetch(`/rota?from=${encodeURIComponent(from)}&to=${encodeURIComponent(to)}`);
  const data = await res.json();

  if (rotaLayer) mapa.removeLayer(rotaLayer);

  if (data.rota) {
    const rotaCoordenadas = data.rota.map(ponto => pontos[ponto]);
    rotaLayer = L.polyline(rotaCoordenadas, { color: "blue", weight: 5 }).addTo(mapa);
    mapa.fitBounds(rotaLayer.getBounds());
  } else {
    alert("Rota não encontrada.");
  }
}
