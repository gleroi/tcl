import * as tcl from "./tcl";

export class Store {
    listeners: ((s :Store) => void)[] = [];
    passages: tcl.PassageArret[] = [];
    favorites: Map<string, boolean> = new Map();

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

    isFavorite(arret: tcl.PassageArret) : boolean {
        return this.favorites.get(arret.id);
    }

    toggleFavorite(arret: tcl.PassageArret) {
        if (!this.favorites.has(arret.id)) {
            this.favorites.set(arret.id, true);
            console.log("toggleFavorite", arret.id, true);
        } else {
            this.favorites.set(arret.id, !this.favorites.get(arret.id))
            console.log("toggleFavorite", arret.id, this.favorites.get(arret.id));
        }
        this.raise();
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