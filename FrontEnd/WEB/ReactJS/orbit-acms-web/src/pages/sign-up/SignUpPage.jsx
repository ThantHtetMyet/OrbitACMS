import { Link } from 'react-router-dom'
import LiquidGlass from 'liquid-glass-react'
import BrandLogo from '../../components/brand-logo/BrandLogo'
import '../sign-in/SignInPage.css'
import './SignUpPage.css'

function SignUpPage() {
  return (
    <section className="auth-page auth-page--signup">
      <span className="auth-orb auth-orb--one" />
      <span className="auth-orb auth-orb--two" />
      <span className="auth-orb auth-orb--three" />

      <LiquidGlass
        className="auth-liquid auth-liquid--signup"
        padding="0"
        cornerRadius={22}
        displacementScale={50}
        blurAmount={0.075}
        saturation={145}
        aberrationIntensity={1.2}
        elasticity={0}
        mode="standard"
        style={{ position: 'absolute', top: '50%', left: '50%' }}
      >
        <form className="auth-card auth-card--signup">
          <div className="auth-header brand-stack">
            <BrandLogo text="entral-ACMS" size="lg" />
          </div>

          <div className="form-grid">
            <div className="field">
              <label htmlFor="firstName">First Name</label>
              <input id="firstName" type="text" placeholder="Enter your first name" />
            </div>
            <div className="field">
              <label htmlFor="lastName">Last Name</label>
              <input id="lastName" type="text" placeholder="Enter your last name" />
            </div>
            <div className="field">
              <label htmlFor="email">Email</label>
              <input id="email" type="email" placeholder="you@example.com" />
            </div>
            <div className="field">
              <label htmlFor="phoneNo">Phone Number</label>
              <input id="phoneNo" type="tel" placeholder="Enter your phone number" />
            </div>
            <div className="field">
              <label htmlFor="password">Password</label>
              <input id="password" type="password" placeholder="Create password" />
            </div>
            <div className="field">
              <label htmlFor="confirmPassword">Confirm Password</label>
              <input id="confirmPassword" type="password" placeholder="Confirm your password" />
            </div>
          </div>

          <button type="button" className="btn btn-primary">Sign Up</button>

          <div className="links links-centered">
            <span className="muted-text">Already have an account?</span>
            <Link className="link" to="/sign-in">Sign In</Link>
          </div>
        </form>
      </LiquidGlass>
    </section>
  )
}

export default SignUpPage
