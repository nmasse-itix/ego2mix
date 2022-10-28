package eco2mix

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchNationalRealTimeData(t *testing.T) {
	expected := `{
		"nhits": 14496,
		"parameters": {
		  "dataset": "eco2mix-national-tr",
		  "rows": 3,
		  "start": 0,
		  "facet": [
			"nature",
			"date_heure"
		  ],
		  "format": "json",
		  "timezone": "UTC"
		},
		"records": [
		  {
			"datasetid": "eco2mix-national-tr",
			"recordid": "19d24188796afae79904f3e38bda937415a7f65b",
			"fields": {
			  "consommation": 52770,
			  "fioul_cogen": 33,
			  "eolien": 977,
			  "bioenergies_dechets": 149,
			  "ech_physiques": 4196,
			  "fioul_tac": 347,
			  "charbon": 28,
			  "bioenergies": 752,
			  "nucleaire": 28172,
			  "prevision_j": 53000,
			  "gaz_ccg": 4063,
			  "hydraulique_step_turbinage": 960,
			  "date": "2022-07-12",
			  "gaz": 4614,
			  "solaire": 8223,
			  "stockage_batterie": "ND",
			  "hydraulique": 5316,
			  "fioul_autres": 113,
			  "heure": "16:45",
			  "taux_co2": 57,
			  "gaz_cogen": 553,
			  "destockage_batterie": "ND",
			  "nature": "Données temps réel",
			  "date_heure": "2022-07-12T14:45:00+00:00",
			  "pompage": -1,
			  "bioenergies_biomasse": 329,
			  "bioenergies_biogaz": 275,
			  "fioul": 493,
			  "eolien_offshore": "ND",
			  "prevision_j1": 52200,
			  "gaz_tac": 0,
			  "gaz_autres": 0,
			  "hydraulique_lacs": 1831,
			  "eolien_terrestre": "ND",
			  "perimetre": "France",
			  "hydraulique_fil_eau_eclusee": 2525
			},
			"record_timestamp": "2022-10-28T19:00:00.369Z"
		  },
		  {
			"datasetid": "eco2mix-national-tr",
			"recordid": "072df8b1f6df8c9fa3890becf14abc9a75cd0417",
			"fields": {
			  "consommation": 52023,
			  "fioul_cogen": 36,
			  "eolien": 1101,
			  "bioenergies_dechets": 153,
			  "ech_physiques": 3968,
			  "fioul_tac": 410,
			  "charbon": 29,
			  "bioenergies": 758,
			  "nucleaire": 28143,
			  "prevision_j": 52250,
			  "gaz_ccg": 4295,
			  "hydraulique_step_turbinage": 1225,
			  "date": "2022-07-12",
			  "gaz": 4816,
			  "solaire": 6393,
			  "stockage_batterie": "ND",
			  "hydraulique": 6256,
			  "fioul_autres": 113,
			  "heure": "17:45",
			  "taux_co2": 60,
			  "gaz_cogen": 522,
			  "destockage_batterie": "ND",
			  "nature": "Données temps réel",
			  "date_heure": "2022-07-12T15:45:00+00:00",
			  "pompage": 0,
			  "bioenergies_biomasse": 329,
			  "bioenergies_biogaz": 275,
			  "fioul": 559,
			  "eolien_offshore": "ND",
			  "prevision_j1": 51450,
			  "gaz_tac": 0,
			  "gaz_autres": 0,
			  "hydraulique_lacs": 2074,
			  "eolien_terrestre": "ND",
			  "perimetre": "France",
			  "hydraulique_fil_eau_eclusee": 2957
			},
			"record_timestamp": "2022-10-28T19:00:00.369Z"
		  },
		  {
			"datasetid": "eco2mix-national-tr",
			"recordid": "98ee9cdc438fd4e2ead5907025a36d210cf22cc1",
			"fields": {
			  "ech_comm_suisse": 1193,
			  "consommation": 52387,
			  "fioul_cogen": 36,
			  "eolien": 1082,
			  "bioenergies_dechets": 154,
			  "ech_physiques": 6782,
			  "fioul_tac": 484,
			  "charbon": 29,
			  "bioenergies": 757,
			  "nucleaire": 28158,
			  "prevision_j": 52400,
			  "ech_comm_angleterre": "2585",
			  "gaz_ccg": 4327,
			  "ech_comm_allemagne_belgique": "3886",
			  "hydraulique_step_turbinage": 1695,
			  "date": "2022-07-12",
			  "gaz": 4852,
			  "solaire": 2561,
			  "stockage_batterie": "ND",
			  "hydraulique": 7569,
			  "fioul_autres": 113,
			  "heure": "19:30",
			  "taux_co2": 65,
			  "gaz_cogen": 526,
			  "destockage_batterie": "ND",
			  "nature": "Données temps réel",
			  "date_heure": "2022-07-12T17:30:00+00:00",
			  "pompage": -36,
			  "bioenergies_biomasse": 330,
			  "bioenergies_biogaz": 275,
			  "fioul": 633,
			  "eolien_offshore": "ND",
			  "prevision_j1": 51600,
			  "ech_comm_italie": -1565,
			  "gaz_tac": 0,
			  "ech_comm_espagne": 900,
			  "gaz_autres": 0,
			  "hydraulique_lacs": 2459,
			  "eolien_terrestre": "ND",
			  "perimetre": "France",
			  "hydraulique_fil_eau_eclusee": 3415
			},
			"record_timestamp": "2022-10-28T19:00:00.369Z"
		  }
		],
		"facet_groups": [
		  {
			"name": "nature",
			"facets": [
			  {
				"name": "Données temps réel",
				"count": 14496,
				"state": "displayed",
				"path": "Données temps réel"
			  }
			]
		  },
		  {
			"name": "date_heure",
			"facets": [
			  {
				"name": "2022",
				"count": 14496,
				"state": "displayed",
				"path": "2022"
			  }
			]
		  }
		]
	  }`

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer svr.Close()

	c := NewEco2mixClient(svr.URL, nil)
	res, err := c.FetchNationalRealTimeData(3)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	if len(res) != 3 {
		t.Errorf("expected 3 records, got %d", len(res))
	}

	expectedTauxCo2 := []int64{57, 60, 65}
	for i := range expectedTauxCo2 {
		if res[i].TauxCo2 != expectedTauxCo2[i] {
			t.Errorf("record %d: expected value %d, got %d", i, expectedTauxCo2[i], res[i].TauxCo2)

		}
	}
}
