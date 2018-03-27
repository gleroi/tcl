import * as React from "react";
import * as tcl from "./tcl";
import { Store } from "./store";

interface AppState {
    arrets: tcl.PassageArret[];
    searchLigne: string;
}

var store = new Store();

export default class App extends React.Component<any, AppState> {
    constructor(props: void, context: any) {
        super(props, context)
        this.state = {
            arrets: [],
            searchLigne: ""
        }
    }

    componentDidMount() {
        store.addChangeListener(s => {
            this.setState(update(this.state, {
                arrets: s.passages
            }));
        });
        store.getPassages("C13A");
    }

    render() {
        return (
            <div>
                <header className="app-header">
                    <h1>TCL</h1>
                </header>
                <div>
                    <form className="search-arret">
                        <label>Ligne:
                            <input type="text" value={this.state.searchLigne} onChange={(e) => this.handleSearchChange(e)} />
                        </label>
                        <button type="button">Chercher</button>
                    </form>
                </div>
                <div className="arrets-container">
                    {this.state.arrets.map((arret: tcl.PassageArret) => (
                        <div key={arret.gid} className="arrets-item">
                        <span className="arrets-footer">
                            <span className="arrets-item-arret">{arret.nom}</span>
                            <button className="arrets-item-favorite"></button>
                        </span>
                            <span className="arrets-item-delai">{arret.delaipassage}</span>
                            <span className="arrets-footer">
                                <span className="arrets-item-ligne">{arret.ligne}</span>
                                <span className="arrets-item-direction">{arret.direction}</span>
                            </span>
                        </div>
                    ))}
                </div>
            </div>
        );
    }

    handleSearchChange(e: React.ChangeEvent<HTMLInputElement>) {
        this.setState(update(this.state, { searchLigne: e.target.value }))
        store.getPassages(e.target.value);
    }

}

function update<T>(original: T, delta: Partial<T>): T {
    return Object.assign({}, original, delta);
}

