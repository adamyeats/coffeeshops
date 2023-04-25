import React from 'react';
import { useQuery } from 'urql';
import Map, {
  FullscreenControl,
  Marker,
  NavigationControl,
  ViewState,
  ViewStateChangeEvent
} from 'react-map-gl';

import { graphql } from '../gql';

import 'mapbox-gl/dist/mapbox-gl.css';
import './MapSheet.css';

const coffeeshopsQuery = graphql(`
  query GetAllCoffeeshops {
    coffeeshops {
      name
      address
      latitude
      longitude
    }
  }
`);

export const MapSheet: React.FC = () => {
  const [{ data }] = useQuery({
    query: coffeeshopsQuery
  });

  const [viewport, setViewport] = React.useState<Partial<ViewState>>({
    latitude: 52.37326665514763,
    longitude: 4.893119196354462,
    zoom: 14
  });

  const onMove = (event: ViewStateChangeEvent) => setViewport(event.viewState);

  return (
    <section className='MapSheet'>
      <Map
        initialViewState={viewport}
        onMove={onMove}
        style={{
          height: '100vh',
          width: '100%'
        }}
        mapStyle='mapbox://styles/mapbox/dark-v11'
        mapboxAccessToken={process.env.REACT_APP_MAPBOX_ACCESS_TOKEN}
      >
        <FullscreenControl position='top-right' />
        <NavigationControl position='top-right' />

        {data?.coffeeshops?.map(coffeeshop => (
          <Marker
            longitude={coffeeshop?.longitude}
            latitude={coffeeshop?.latitude}
            anchor='bottom'
          />
        ))}
      </Map>
    </section>
  );
};
