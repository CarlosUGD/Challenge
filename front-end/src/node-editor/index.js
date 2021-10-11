import Rete from "rete";
import VueRenderPlugin from "rete-vue-render-plugin";
import ConnectionPlugin from "rete-connection-plugin";
import AreaPlugin from "rete-area-plugin";
import ContextMenuPlugin from "rete-context-menu-plugin";
import { NumComponent } from "./components/numComponent";
import { AddComponent } from "./components/addComponent";
import { SubComponent } from "./components/subComponent";
import { MultComponent } from "./components/multComponent";
import { DivComponent } from "./components/divComponent";

export default async function(container) {
  var components = [
    new NumComponent(),
    new AddComponent(),
    new SubComponent(),
    new MultComponent(),
    new DivComponent(),
  ];

  var editor = new Rete.NodeEditor("demo@0.1.0", container);
  editor.use(ConnectionPlugin);
  editor.use(VueRenderPlugin);
  editor.use(ContextMenuPlugin);
  editor.use(AreaPlugin);

  var engine = new Rete.Engine("demo@0.1.0");

  components.map((c) => {
    editor.register(c);
    engine.register(c);
  });

  editor.on("process nodecreated noderemoved connectioncreated connectionremoved", async () => {
    await engine.abort();
    await engine.process(editor.toJSON());
  });

  editor.view.resize();
  AreaPlugin.zoomAt(editor);
  editor.trigger("process");
}
