// simulo chiamata API
function getDataFromAPI(qrcontent){
    fetch('/api/operazioni',{ method:'post',body: JSON.stringify({content: qrcontent}) })
        .then(r=> {
            r.json().then(risp => {
                console.log(risp)
                document.querySelector("#output").appendChild(createTableFromJson(risp));
            })
        })
        .catch(alert)
}


function createTableFromJson(input) {
    if(input != null && input.length > 0) {
        const headers = Object.keys((input[0]))

        console.log(headers)
        const table = document.createElement('table');
        table.border = '1';

        // Create a table headers
        const rowHeaders = document.createElement('tr');
        headers.forEach(header => {
            const cell = document.createElement('td');
            cell.textContent = header;
            rowHeaders.appendChild(cell);
        });
        // Append the row to the table
        table.appendChild(rowHeaders);

        // Iterate over the array of arrays
        input.forEach(rowData => {
            // Create a table row element
            const row = document.createElement('tr');

            headers.forEach(header => {
                const cell = document.createElement('td');
                cell.textContent = rowData[header];
                row.appendChild(cell);
            });

            // Append the row to the table
            table.appendChild(row);
        });

        return table
    } else {
        const msg = document.createElement('p')
        msg.textContent = "Nessuna operazione trovata"
        return msg
    }
}