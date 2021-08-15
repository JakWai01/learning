use yew::prelude::*;
use web_sys::FocusEvent;
use std::{collections::HashMap};

enum Msg {
    Clicked,
}

struct Model {
    // 'ComponentLink' is like a reference to a component. 
    // It can be used to send messages to the component
    link: ComponentLink<Self>,
    value: i64,
}

impl Component for Model {
    type Message = Msg;
    type Properties = ();

    fn create(_props: Self::Properties, link: ComponentLink<Self>) -> Self {
        Self {
            link,
            value: 0,
        }
    }

    fn update(&mut self, msg: Self::Message) -> ShouldRender {
        match msg {
            Msg::Clicked => {
                self.value += 1;
                true
            }
        }
    }

    fn change(&mut self, _props: Self::Properties) -> ShouldRender {
        // Should only return "true" if new properties are different to
        // previously received properties.ComponentLink
        // This component has no properties so we will always return "false".
        false
    }

    fn view(&self) -> Html {
        
        html! {
            <div class=classes!("h-screen", "flex", "justify-center")>
                <div class=classes!("h-48", "m-auto", "flex", "items-center")>
                    <form onsubmit={self.link.callback(|e: FocusEvent| {e.prevent_default(); Msg::Clicked})}>
                        <label for="name" class=classes!("text-3xl", "inline-block", "w-48", "mr-6", "text-right", "font-bold", "text-gray-600")>{"How bad in"}</label>
                        <input placeholder="e.g. New York" class=classes!("mr-6", "text-3xl", "flex-1", "py-2", "border-b-2", "border-gray-400", "focus:border-green-400", "text-gray-600", "placeholder-gray-400", "outline-none") type="text" id="fname" name="fname"/>
                        <label class=classes!("text-3xl")>{"?"}</label>
                    </form>
                </div>
                <h1>{self.value}</h1>
            </div>
        }
    }
}

// How to get the coordinates for a location
// http://api.openweathermap.org/geo/1.0/direct?q=London&limit=5&appid={API key}
// How to get the elevation?
// https://github.com/Jorl17/open-elevation/blob/master/docs/api.md#

#[tokio::main]
async fn main() {
    let data = get_coordinates().await.unwrap(); 

    for (key, value) in &data {
        println!("{:?} -> {}", key, value);
    }

    yew::start_app::<Model>();
}

async fn get_coordinates() -> Result<HashMap<String, String>, Box<dyn std::error::Error>> {
    let resp = reqwest_wasm::get("http://api.positionstack.com/v1/forward?access_key=425ffa7e6a357e019092c097563e15fd&query=%20Stuttgart").await?.json::<HashMap<String, String>>().await?;
    println!("{:?}", resp);
    Ok(resp)
}