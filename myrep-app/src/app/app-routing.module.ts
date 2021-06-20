import { Routes, RouterModule } from '@angular/router';

import { HomeComponent } from './home';
import { ProfileComponent } from './profile';
import { AuthGuard } from '@auth0/auth0-angular';
import { UserComponent } from './user/user.component';

const routes: Routes = [
    { path: '', component: HomeComponent },
    { path: 'profile', component: ProfileComponent, canActivate: [AuthGuard]},
    { path: 'users/:username', component: UserComponent},

    // otherwise redirect to home
    { path: '**', redirectTo: '' }
];

export const AppRoutingModule = RouterModule.forRoot(routes);