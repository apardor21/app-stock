let stocksTable;
let marketTable;

$(document).ready(function () {

  // ===============================
  // Inicializar Materialize
  // ===============================
  $('select').formSelect();

  // ===============================
  // DATA TABLE: STOCKS
  // ===============================
  stocksTable = $('#stocksTable').DataTable({
    ajax: {
      url: 'http://localhost:8080/api/stocks',
      dataSrc: ''
    },
    columns: [
      { data: 'symbol' },
      { data: 'price' },
      { data: 'change_percent' },
      { data: 'volume' },
      { data: 'source' },
      { data: 'created_at' }
    ],
    pageLength: 5,
    lengthChange: false
  });

  // Botón para consumir Alpha Vantage (Stocks)
  $('#btnFetch').click(function () {
    const symbol = $('#symbolSelect').val();
    fetchStock(symbol);
  });

  // ===============================
  // DATA TABLE: MARKET STATUS
  // ===============================
  marketTable = $('#marketStatusTable').DataTable({
  ajax: {
    url: 'http://localhost:8080/api/market-status',
    dataSrc: function (json) {

      // Si backend devuelve error
      if (json.error) {
        M.toast({ html: json.error });
        return []; //  DataTables SIEMPRE espera array
      }

      // Si no es array (seguridad extra)
      if (!Array.isArray(json)) {
        return [];
      }

      return json;
    }
  },
  columns: [
    { data: 'market_type' },
    { data: 'region' },
    { data: 'primary_exchanges' },
    {
      data: 'current_status',
      render: function (data, type) {
        if (type === 'filter' || type === 'sort') {
          return data;
        }
        return data === 'open'
          ? '<span class="green-text"><b>OPEN</b></span>'
          : '<span class="red-text"><b>CLOSED</b></span>';
      }
    },
    { data: 'local_open' },
    { data: 'local_close' }
  ],
  pageLength: 5,
  lengthChange: false
});


  // ===============================
  // Auto-refresh Market Status
  // ===============================
  setInterval(() => {
    marketTable.ajax.reload(null, false);
  }, 60000); // cada 1 minuto
});

// ===============================
// Función: Fetch Stock desde API
// ===============================
function fetchStock(symbol) {

  if (!symbol) {
    M.toast({ html: 'Selecciona un símbolo' });
    return;
  }

  fetch(`http://localhost:8080/api/stocks/fetch?symbol=${symbol}`)
    .then(res => res.json())
    .then(data => {
      M.toast({ html: 'Stock actualizado correctamente' });
      stocksTable.ajax.reload(null, false);
    })
    .catch(err => {
      console.error(err);
      M.toast({ html: 'Error al obtener el stock' });
    });
}
