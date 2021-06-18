import { Routes, RouterModule } from '@angular/router';

import { HomeComponent } from './home';

const routes: Routes = [
    { path: '', component: HomeComponent },

    // otherwise redirect to home
    { path: '**', redirectTo: '' }
];

export const AppRoutingModule = RouterModule.forRoot(routes);