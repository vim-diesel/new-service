import { useUser } from '@clerk/clerk-react';

export default function Home() {
  const { isSignedIn, user, isLoaded } = useUser();

  if (!isLoaded) {
    return null;
  }

  if (isSignedIn) {
    return (
      <>
        <div>Hello {user.firstName}!</div>
      </>
    );
  }

  return <div>Not signed in</div>;
}
