import React, {Component} from 'react';
import {Route, Switch} from 'react-router-dom';
import {NavigationDrawer} from 'react-md';
import NavItemLink from './NavItemLink';
import AddSite from './AddSite'
import ListSites from './ListSites'
import './App.css';

const navItems = [{
    label: 'Add Site',
    to: '/',
    exact: true,
    icon: 'inbox',
}, {
    label: 'List Sites',
    to: '/sites',
    icon: 'star',
}];

class App extends Component {
    render() {
        return (
            <NavigationDrawer
                drawerTitle="KeyWee"
                toolbarTitle="Facebook Interest Extractor"
                navItems={navItems.map(props => <NavItemLink {...props} key={props.to}/>)}
            >
                <Switch>
                    <Route path='/' exact component={AddSite}/>
                    <Route path='/sites' component={ListSites}/>
                </Switch>
            </NavigationDrawer>
        );
    }
}

export default App;
