import React from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter} from 'react-router-dom'
import './index.css';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
import WebFontLoader from 'webfontloader';

WebFontLoader.load({
    google: {
        families: ['Roboto:300,400,500,700', 'Material Icons'],
    },
});
ReactDOM.render(
    <BrowserRouter>
        <App/>
    </BrowserRouter>
    , document.getElementById('root')
);
registerServiceWorker();
