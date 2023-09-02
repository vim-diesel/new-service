import './App.css'
import ProfileCard from './components/MyCard'


function App() {
  return (
    <>
      <div className="flex flex-row ">
        <div className="flex-basis-1/3 grow">
          <ProfileCard
            title="My Profile"
            description="This is my profile"
            content="This is my content"
            footer="This is my footer"
            image="../public/cheng-feng-j7AMlh2MMHc-unsplash.jpg" />
        </div>
        <div className="flex-basis-1/3 grow">
          <ProfileCard
            title="My Profile"
            description="This is my profile"
            content="This is my content"
            footer="This is my footer"
            image="../public/kevin-laminto-7PqRZK6rbaE-unsplash.jpg" />
        </div>
        <div className="flex-basis-1/3 grow">
          <ProfileCard
            title="My Profile"
            description="This is my profile"
            content="This is my content"
            footer="This is my footer"
            image="../public/levon-vardanyan-ihUVzI4f5To-unsplash.jpg" />
        </div>
      </div >
    </>
  )
}

export default App
