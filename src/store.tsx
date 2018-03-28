import * as tcl from "./tcl";

export class Store {
    listeners: ((s :Store) => void)[] = [];
    passages: tcl.PassageArret[] = [];
    favorites: Map<string, boolean> = new Map();

    constructor() {
        this.loadFavorites();
    }

    getPassages(ligne: string) {
        tcl.getPassageArrets(ligne)
        .then(arrets => {
            var groups = new Map<string, tcl.PassageArret>();
            arrets.map(arret => {
                let key = arret.id + arret.direction;
                if (groups.has(key)) {
                    let other = groups.get(key);
                    this.mergeArret(other, arret)
                    groups.set(key, other);
                } else {
                    arret.delais = [arret.delaipassage, "NA"]
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

    mergeArret(a1 : tcl.PassageArret, a2: tcl.PassageArret) {
        if (this.delai(a1) < this.delai(a2)) {
            a1.delais = [a1.delaipassage, a2.delaipassage];
        } else {
            a1.delais = [a2.delaipassage, a1.delaipassage];
        }
    }

    delai(a : tcl.PassageArret) : number {
        return parseInt(a.delaipassage);
    }

    isFavorite(arret: tcl.PassageArret) : boolean {
        return this.favorites.get(arret.id);
    }

    toggleFavorite(arret: tcl.PassageArret) {
        if (!this.favorites.has(arret.id)) {
            this.favorites.set(arret.id, true);
        } else {
            this.favorites.delete(arret.id);
        }
        this.saveFavorites()
        this.raise();
    }

    saveFavorites() {
        let entries = Array.from(this.favorites.entries())
        localStorage.setItem("favorites", JSON.stringify(entries))
    }

    loadFavorites() {
        let json = localStorage.getItem("favorites")
        let entries : [string, boolean][] = JSON.parse(json);
        this.favorites = new Map(entries);
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