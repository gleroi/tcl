import * as React from "react";
import * as tcl from "./tcl";
import { Store } from "./store";

interface AppState {
    arrets: tcl.PassageArret[];
    favorites: tcl.PassageArret[];
    searchLigne: string;
    searchEnable: boolean;
}

var store = new Store();

export default class App extends React.Component<any, AppState> {
    constructor(props: void, context: any) {
        super(props, context)
        this.state = {
            arrets: [],
            favorites: [],
            searchLigne: "",
            searchEnable: false
        }
    }

    componentDidMount() {
        store.addChangeListener(s => {
            this.setState(update(this.state, {
                arrets: s.passages,
                favorites: s.Favorites()
            }));
        });
        store.raise();
    }

    render() {
        return (
            <div>
                <header className="app-header">
                    <h1>TCL</h1>
                    <button className="toggle-search" type="button" onClick={(e) => this.toggleSearch()}>
                        <span className="fa fa-search"></span>
                    </button>
                </header>
                <div className={"search-arret " + (this.state.searchEnable ? "" : "hidden")}>
                    <label>Ligne:
                        <input type="text" value={this.state.searchLigne} onChange={(e) => this.handleSearchChange(e)} />
                    </label>
                </div>
                <PassageList arrets={this.state.searchEnable ? this.state.arrets : this.state.favorites} />
            </div>
        );
    }

    handleSearchChange(e: React.ChangeEvent<HTMLInputElement>) {
        this.setState(update(this.state, { searchLigne: e.target.value }))
        store.getPassages(e.target.value);
    }


    toggleSearch() {
        this.setState(update(this.state, {
            searchEnable: !this.state.searchEnable
        }));
    }

}

interface PassageListProps {
    arrets: tcl.PassageArret[];
}

class PassageList extends React.Component<PassageListProps, any> {
    render() {
        console.log(this.props.arrets)
        var content = [<h1 key="banner-1" className="arrets-help-banner">
            Saisir le nom d'un ligne et mettre en favori les arrets
        </h1>];
        if (this.props.arrets.length > 0) {
            content = this.props.arrets.map((arret: tcl.PassageArret) => (
                <div key={arret.gid} className="arrets-item" onClick={(e) => this.toggleFavorite(arret)}>
                    <span className="arrets-footer">
                        <span className="arrets-item-arret">{arret.nom}</span>
                        {store.isFavorite(arret) && (
                            <span className="arrets-item-favorite"></span>
                        )}
                    </span>
                    <span className="arrets-item-delai">{arret.delais[0]}</span>
                    <span className="arrets-item-delai">{arret.delais[1]}</span>
                    <span className="arrets-footer">
                        <span className="arrets-item-ligne">{arret.ligne}</span>
                        <span className="arrets-item-direction">{arret.direction}</span>
                    </span>
                </div>
            ));
        }
        return (
            <div className="arrets-container">
                {content}
            </div>
        );
    }

    
    toggleFavorite(arret: tcl.PassageArret) {
        store.toggleFavorite(arret);
    }
}

function update<T>(original: T, delta: Partial<T>): T {
    return Object.assign({}, original, delta);
}

