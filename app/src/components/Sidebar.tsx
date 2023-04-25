import React from 'react';
import './Sidebar.css';

export const Sidebar: React.FC = () => (
  <aside className='Sidebar'>
    <h1 className='Sidebar__title'>
      These are the <span>Coffeeshops</span> of Amsterdam.
    </h1>
    <p>
      Coffeeshops in Amsterdam have a rich history dating back to the 1970s. The
      Dutch government adopted a policy of tolerance towards soft drugs, which
      allowed for the establishment of coffeeshops as places where people could
      buy and consume cannabis without fear of prosecution. These coffeeshops
      quickly became popular among locals and tourists alike, and Amsterdam soon
      became known as a haven for cannabis enthusiasts.
    </p>
    <p>
      In the early years, coffeeshops in Amsterdam were often small and dingy,
      with a limited selection of products. However, as the cannabis culture
      grew, so did the coffeeshops. Today, there are over 150 coffeeshops in
      Amsterdam. Coffeeshops have become an important part of Amsterdam's
      culture and economy, attracting millions of visitors each year.
    </p>
    <p>
      This small distraction was made by <a href='http://adamyeats.dev'>this person.</a>
    </p>
  </aside>
);
