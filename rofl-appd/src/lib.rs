//! REST API daemon accessible by ROFL apps.

mod routes;
pub mod services;
pub(crate) mod state;
pub mod types;

use std::sync::Arc;

use rocket::{figment::Figment, routes};

use oasis_runtime_sdk::modules::rofl::app::{App, Environment};

/// API server configuration.
#[derive(Clone)]
pub struct Config<'a> {
    /// Address where the service should listen on.
    pub address: &'a str,
    /// Key management service to use.
    pub kms: Arc<dyn services::kms::KmsService>,
}

/// Start the REST API server.
pub async fn start<A>(cfg: Config<'_>, env: Environment<A>) -> Result<(), rocket::Error>
where
    A: App,
{
    // Oasis runtime environment.
    let env: Arc<dyn state::Env> = Arc::new(state::EnvImpl::new(env));

    // Server configuration.
    let rocket_cfg = Figment::new().join(("address", cfg.address));

    rocket::custom(rocket_cfg)
        .manage(env)
        .manage(cfg.kms)
        .mount("/rofl/v1/app", routes![routes::app::id,])
        .mount("/rofl/v1/keys", routes![routes::keys::generate,])
        .launch()
        .await?;

    Ok(())
}