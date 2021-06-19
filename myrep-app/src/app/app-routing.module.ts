import { Routes, RouterModule } from '@angular/router';

import { HomeComponent } from './home';
import { ProfileComponent } from './profile';

const routes: Routes = [
    { path: '', component: HomeComponent },
    { path: 'profile', component: ProfileComponent },

    // otherwise redirect to home
    { path: '**', redirectTo: '' }
];

export const AppRoutingModule = RouterModule.forRoot(routes);