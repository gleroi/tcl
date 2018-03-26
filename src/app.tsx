import * as React from "react";
import * as tcl from "./tcl";

interface AppState {
    arrets: tcl.PassageArret[];
    searchLigne: string;
}

export default class App extends React.Component<any, AppState> {
    constructor(props: void, context: any) {
        super(props, context)
        this.state = {
            arrets: [],
            searchLigne: ""
        }
    }

    componentDidMount() {
        tcl.getPassageArrets("C13A")
            .then(stops => this.setState({ arrets: stops }))
            .catch(err => console.error(err));
    }

    render() {
        return (
            <div>
                <h1>TCL</h1>
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
                        <div className="arrets-item">
                            <span className="arrets-item-arret">{arret.idtarretdestination}</span>
                            <span className="arrets-item-delai">{arret.delaipassage}</span>
                            <span className="arrets-item-ligne">{arret.ligne}</span>
                        </div>
                    ))}
                </div>
            </div>
        );
    }

    handleSearchChange(e: React.ChangeEvent<HTMLInputElement>) {
        this.setState(update(this.state, { searchLigne: e.target.value }))
        tcl.getPassageArrets(e.target.value)
            .then(stops => this.setState({ arrets: stops }))
            .catch(err => console.error(err));
    }

}

function update<T>(original: T, delta: Partial<T>) : T {
    return Object.assign({}, original, delta);
}

