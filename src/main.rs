use yew::prelude::*;
use web_sys::FocusEvent;

enum Msg {
    AddOne,
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
            Msg::AddOne => {
                self.value += 1;
                //the value has changed so we need to 
                // re-render for it to appear on the page
                true
            },
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




fn main() {
    yew::start_app::<Model>();
}
