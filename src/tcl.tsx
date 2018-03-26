import {Base64 as base64} from "js-base64";

export interface PassageArret {
    idtarretdestination: string;
    coursetheorique: string;
    direction: string;
    ligne: string;
    delaipassage: string;
    heurepassage: string;
    gid: string;
    last_update_fme: string;
    type: string;
    id: string;
}

export function getPassageArrets(ligne: string): Promise<PassageArret[]> {
    const url = "/api/passagearret?compact=true";

    return fetch(url + "&field=ligne&value=" + ligne, {
        method: "GET",
        headers: new Headers({
            'Accept': 'application/json',
        })
    })
    .then(r => r.json())
    .then(json => {
        if (json == null || json.values == null) {
            throw new Error("passage aux arrets: ligne " + ligne + ": reponses vide");
        }
        return json.values;
    });
}