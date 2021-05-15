#![feature(box_patterns, proc_macro_diagnostic)]
#![deny(rust_2018_idioms)]

use proc_macro::TokenStream;

mod error_derive;
mod event_derive;
mod generators;
mod handlers_attr;
#[cfg(test)]
mod test_utils;
mod version_from_cargo;

/// Derives the `Event` trait on an enum.
#[proc_macro_derive(Event, attributes(sdk_event))]
pub fn event_derive(input: TokenStream) -> TokenStream {
    let input = syn::parse_macro_input!(input as syn::DeriveInput);
    event_derive::derive_event(input).into()
}

/// Derives the `Error` trait on an enum.
// The helper attribute is `sdk_error` to avoid conflict with `thiserror::Error`.
#[proc_macro_derive(Error, attributes(sdk_error, source, from))]
pub fn error_derive(input: TokenStream) -> TokenStream {
    let input = syn::parse_macro_input!(input as syn::DeriveInput);
    error_derive::derive_error(input).into()
}

/// Constructs an `oasis_sdk::core::common::version::Version` from the Cargo.toml version.
#[proc_macro]
pub fn version_from_cargo(_input: TokenStream) -> TokenStream {
    version_from_cargo::version_from_cargo().into()
}

/// "Derives" `MethodRegistrationHandler` from an `impl` item.
#[proc_macro_attribute]
pub fn handlers(args: TokenStream, input: TokenStream) -> TokenStream {
    let input = syn::parse_macro_input!(input as syn::ItemImpl);
    let args = syn::parse_macro_input!(args as syn::AttributeArgs);
    handlers_attr::gen_method_registration_handler_impl(input, args).into()
}

/// A "helper attribute" for the handlers "derive." This attribute could be stripped
/// by the `handlers` attribute, but if it's accidentally omitted, not having this
/// one will give really confusing error messages.
#[proc_macro_attribute]
pub fn handler(_args: TokenStream, input: TokenStream) -> TokenStream {
    // `sdk::handler` can only be applied to methods, of course.
    let input = syn::parse_macro_input!(input as syn::ImplItemMethod);
    quote::quote!(#input).into()
}
