declare module "leaflet" {
  export type LatLngTuple = [number, number];

  export interface DivIconOptions {
    className?: string;
    html?: string;
    iconSize?: LatLngTuple;
    iconAnchor?: LatLngTuple;
  }

  export type DivIcon = object;
  export type LatLngBounds = object;

  const L: {
    divIcon(options: DivIconOptions): DivIcon;
    latLngBounds(latlngs: LatLngTuple[]): LatLngBounds;
  };

  export default L;
}
