import React, {Component} from 'react';
import {ExpansionList, ExpansionPanel} from 'react-md';
import SitesAPI from "./api";

class ListSites extends Component {
    state = {
        sites: {},
        toasts: [],
        autohide: true
    };

    dismissToast = () => {
        const [, ...toasts] = this.state.toasts;
        this.setState({toasts});
    };

    componentDidMount() {
        this._asyncRequest = SitesAPI.all().then(sites => {
            this.setState({
                sites: sites,
            });
        }).catch((err) => {
            this.setState((state) => {
                const toasts = state.toasts.slice();
                toasts.push({text: "Had an error retriving data", action: false});
                return {toasts};
            });
        });
    }

    componentWillUnmount() {
        if (this._asyncRequest) {
            this._asyncRequest = null;
        }
    }

    render() {
        return (
            <ExpansionList className='md-cell md-cell--12'>
                {
                    Object.keys(this.state.sites).map((i) => {
                        let s = this.state.sites[i];
                        return <ExpansionPanel label={s.url} footer={null}>
                            <p>Audiance Size: {s.audianceSize}</p>
                            <p>Estimate Ready: {s.estimateReady}</p>
                            <p>Topics: {s.topicsList.join(", ")}</p>
                            <p>Geo: {s.geoList.join(". ")}</p>
                            <p>Facebook Intrests: <ul>{
                                s.facebookIntrestsList.map((fi) => {
                                    return <li>id: {fi.id} Name: {fi.name}</li>
                                })
                            }</ul>
                            </p>
                        </ExpansionPanel>
                    })
                }
            </ExpansionList>
        );
    }
}

export default ListSites;