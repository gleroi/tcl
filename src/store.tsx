import * as tcl from "./tcl";

export class Store {
    listeners: ((s: Store) => void)[] = [];
    passages: tcl.PassageArret[] = [];
    favorites: Map<string, tcl.PassageArret> = new Map();

    constructor() {
        this.loadFavorites();
        this.reloadFavorites()
    }

    getPassages(ligne: string) {
        tcl.getPassageArrets(ligne)
            .then(arrets => {
                this.passages = this.merge(arrets);
                this.raise();
            })
            .catch(err => {
                console.error("retrieving passages arrets", err);
            });
    }

    getFavorites(): tcl.PassageArret[] {
        var r = Array.from(this.favorites.values());
        return r;
    }

    reloadFavorites() {
        var all = [];
        for (let key of this.favorites.keys()) {
            all.push(
                ((id: string) => {
                    return tcl.getPassageById(id)
                        .then(arrets => {
                            this.favorites.set(id, this.merge(arrets)[0]);
                            this.raise();
                        })
                })(key));
        }
        Promise.all(all).then(results => {
            this.saveFavorites();
        })
    }

    merge(arrets: tcl.PassageArret[]): tcl.PassageArret[] {
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
        return Array.from(groups.values());
    }

    mergeArret(a1: tcl.PassageArret, a2: tcl.PassageArret) {
        if (this.delai(a1) < this.delai(a2)) {
            a1.delais = [a1.delaipassage, a2.delaipassage];
        } else {
            a1.delais = [a2.delaipassage, a1.delaipassage];
        }
    }

    delai(a: tcl.PassageArret): number {
        if (a.delaipassage == "Proche") {
            return 0;
        }
        return parseInt(a.delaipassage);
    }

    isFavorite(arret: tcl.PassageArret): boolean {
        return this.favorites.has(arret.id);
    }

    toggleFavorite(arret: tcl.PassageArret) {
        if (!this.favorites.has(arret.id)) {
            this.favorites.set(arret.id, arret);
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
        let entries: [string, tcl.PassageArret][] = JSON.parse(json);
        this.favorites = new Map(entries);
    }

    raise() {
        for (let fn of this.listeners) {
            fn(this);
        }
    }

    addChangeListener(fn: (s: Store) => void) {
        this.listeners.push(fn);
    }
}