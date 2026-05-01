import { Link } from 'react-router-dom'
import LiquidGlass from 'liquid-glass-react'
import BrandLogo from '../../components/brand-logo/BrandLogo'
import './SignInPage.css'

function SignInPage() {
  return (
    <section className="auth-page auth-page--signin">
      <span className="auth-orb auth-orb--one" />
      <span className="auth-orb auth-orb--two" />
      <span className="auth-orb auth-orb--three" />

      <LiquidGlass
        className="auth-liquid auth-liquid--signin"
        padding="0"
        cornerRadius={22}
        displacementScale={66}
        blurAmount={0.075}
        saturation={145}
        aberrationIntensity={2.5}
        elasticity={0}
        mode="standard"
        style={{ position: 'absolute', top: '50%', left: '50%' }}
      >
        <form className="auth-card auth-card--signin">
          <div className="auth-header brand-stack">
            <BrandLogo text="entral-ACMS" size="lg" />
          </div>

          <div className="field">
            <label htmlFor="userId">UserID</label>
            <input id="userId" type="text" placeholder="Enter your UserID" />
          </div>

          <div className="field">
            <label htmlFor="password">Password</label>
            <input id="password" type="password" placeholder="Enter your password" />
          </div>

          <button type="button" className="btn btn-primary">Sign In</button>

          <div className="links links-split">
            <a className="link" href="#">Forgot password/UserID</a>
            <Link className="link" to="/sign-up">Create account</Link>
          </div>
        </form>
      </LiquidGlass>
    </section>
  )
}

export default SignInPage
