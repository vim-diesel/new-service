import './App.css';
import ProfileCard from './components/cards/MyCard';
import { Separator } from '@/components/ui/separator';
import { Button } from '@/components/ui/button';
import {
  ClerkProvider,
  SignedIn,
  SignedOut,
  RedirectToSignIn,
} from '@clerk/clerk-react';
import Example from './components/example';
import Home from './components/pages/Home';

if (!import.meta.env.VITE_REACT_APP_CLERK_PUBLISHABLE_KEY) {
  throw new Error('Missing Publishable Key');
}
const clerkPubKey = import.meta.env.VITE_REACT_APP_CLERK_PUBLISHABLE_KEY;

function App() {
  return (
    <>
      <ClerkProvider publishableKey={clerkPubKey}>
        <SignedIn>
          <div className='flex flex-row'>
            <div>
              <ProfileCard
                cardTitle='My Profile'
                cardDescription='This is my profile'
                cardContent='This is my content'
                cardFooter='This is my footer'
                cardImage='/cheng-feng-j7AMlh2MMHc-unsplash.jpg'
              />
            </div>
            <div>
              <ProfileCard
                cardTitle='My Profile'
                cardDescription='This is my profile'
                cardContent='This is my content'
                cardFooter='This is my footer'
                cardImage='/kevin-laminto-7PqRZK6rbaE-unsplash.jpg'
              />
            </div>
            <div>
              <ProfileCard
                cardTitle='My Profile'
                cardDescription='This is my profile'
                cardContent='This is my content'
                cardFooter='This is my footer'
                cardImage='/levon-vardanyan-ihUVzI4f5To-unsplash.jpg'
                cardImageAlt='Image Alt Text'
              />
            </div>
          </div>
          <Separator className='my-4' />
          <Button>Add city</Button>
          <Separator className='my-4' />
          <Example />
          <Separator className='my-4' />
          <Home />
        </SignedIn>
        <SignedOut>
          <RedirectToSignIn />
        </SignedOut>
      </ClerkProvider>
    </>
  );
}

export default App;
