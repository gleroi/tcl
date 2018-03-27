import * as tcl from "./tcl";

export class Store {
    listeners: ((s :Store) => void)[] = [];
    passages: tcl.PassageArret[] = [];

    getPassages(ligne: string) {
        tcl.getPassageArrets(ligne)
        .then(arrets => {
            var groups = new Map<string, tcl.PassageArret>();
            arrets.map(arret => {
                let key = arret.id + arret.direction;
                if (groups.has(key)) {
                    let other = groups.get(key);
                    if (arret.delaipassage < other.delaipassage) {
                        groups.set(key, arret);
                    }
                } else {
                    groups.set(key, arret);
                }
            });
            this.passages = Array.from(groups.values());
            this.raise();
        })
        .catch(err => {
            console.error("retrieving passages arrets", err);
        });
    }

    raise() {
        for (let fn of this.listeners) {
            fn(this);
        }
    }

    addChangeListener(fn : (s: Store) => void) {
        this.listeners.push(fn);
    }
}