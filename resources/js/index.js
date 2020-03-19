window.addEventListener("popstate", function (e) {
    history.pushState(null, null, null);
    return;
});

document.querySelector("#search-button").addEventListener("click", renderResults);

async function renderResults() {
  try {
    let terms = document.querySelector("#terms").value;
    let page = document.querySelector("#page").value;
    let responseMapList = await Search(terms, page); // call go function
    let responseKeys = await GetKeysSearchResponse(); // call go function
    console.log("Search results:");
    console.log(responseMapList);
    renderTable(responseKeys, responseMapList)
  } catch (err) {
    alert("Error:" + err);
  }
}


function renderTable(responseKeys, results) {
  let tableElement = createCustomElement("table", null, (e) => {
    e.classList.add("table");
    e.classList.add("text-nowrap");
  });
  let tr = createCustomElement("tr");
  // Table header
  responseKeys.forEach((key) => tr.appendChild(createCustomElement("th", key)));
  let thead = createCustomElement("thead", null, (e) => {
    e.appendChild(tr);
    e.classList.add("thead-light");
  });
  tableElement.appendChild(thead);

  // Table body
  let tbody = createCustomElement("tbody");
  results.forEach((result) => {
    let tr = createCustomElement("tr");
    responseKeys.forEach((key) => tr.appendChild(createCustomElement("td", result[key])));
    tbody.appendChild(tr);
  });
  tableElement.appendChild(tbody);
  
  document.querySelector("#result-container").innerHTML = '';
  document.querySelector("#result-container").appendChild(tableElement);
}

// var formElement = document.getElementById("form-container");
//
// formElement.addEventListener("submit", function(event) {
//     event.preventDefault();
//     return false;
// });
//
// formElement.addEventListener("submit", async () => {
//     var query = document.getElementById("form-query").value;
//     try {
//         var response = await getResponseSearch(query); // call go function
//         console.log(response);
//         renderTable(document.getElementById("result-tbody-container-id"), response.results);
//     } catch (err) {
//         alert("Error:" + err);
//     }
// });
//
// function renderTable(tbodyElement, results) {
//     console.log(tbodyElement);
//     console.log(results);
//     // remove all child node
//     while(tbodyElement.firstChild) {
//         tbodyElement.removeChild(tbodyElement.firstChild);
//     }
//
//     for (index in results) {
//
//         entity = results[index];
//         var tr = document.createElement("tr");
//
//         // artworkUrl100
//         var tdThumbnail = createCustomElement("td", null, function (element) {
//             var img = createCustomElement("img", null, function (element) { element.src = entity.artworkUrl100; });
//             img.style.minWidth = "300px"
//             var audio = createCustomElement("audio", null, function (element) {
//                 element.src = entity.previewUrl;
//                 element.controls = 'controls';
//             });
//             element.style.textAlign = "center"
//             element.appendChild(img);
//             element.appendChild(audio);
//         });
//         tr.appendChild(tdThumbnail);
//
//         // Artist
//         var tdArtist = createCustomElement("td", null, function(element) {
//             var spanIdBadge = createCustomElement("span", "artistId: " + entity.artistId);
//             spanIdBadge.classList.add("badge", "badge-secondary")
//             var pArtistName = createCustomElement("p", entity.artistName);
//             element.appendChild(spanIdBadge);
//             element.appendChild(pArtistName);
//         });
//         tr.appendChild(tdArtist);
//
//         // Track
//         var tdTrack = createCustomElement("td", null, function (element) {
//             var spanIdBadge = createCustomElement("span", "trackId: " + entity.trackId);
//             spanIdBadge.classList.add("badge", "badge-secondary")
//             var pTrackName = createCustomElement("p", entity.trackName);
//             element.appendChild(spanIdBadge);
//             element.appendChild(pTrackName);
//         });
//         tr.appendChild(tdTrack);
//
//         // ReleaseDate
//         var tdReleaseDate = createCustomElement("td", null, function (element) {
//             var pReleaseDate = createCustomElement("p", entity.releaseDate);
//             element.appendChild(pReleaseDate);
//         });
//         tr.appendChild(tdReleaseDate);
//
//         // TrackPrice
//         var tdTrackProce = createCustomElement("td", null, function (element) {
//             var pTrackPrice = createCustomElement("p", entity.trackPrice);
//             element.appendChild(pTrackPrice);
//         });
//         tr.appendChild(tdTrackProce);
//
//
//         tbodyElement.appendChild(tr)
//     }
// }

function createCustomElement(tagName, innerText = null, customProcessing = null) {
  let element = document.createElement(tagName);
  if (innerText != null) {
    element.innerText = innerText;
  }
  if (customProcessing != null) {
    customProcessing(element);
  }
  return element;
}
