import React from 'react';
import {
  Container,
  Media,
  MediaLeft,
  MediaContent,
  MediaRight,
  Image,
  Title,
  Content,
  Button,
} from 're-bulma';

import Icon from '../Icons';

import './style.css';

const ClassifiedCard = ({ organization }) => (
  <Container isFullwidth>
    <Media className="classifiedCard">
      <MediaLeft>
        <Image src={organization.image} alt={`${organization.name}`} />
      </MediaLeft>
      <MediaContent>
        <Content>
          <div className="organizationContent">
            <Title size="is5">{organization.name}</Title>
            <p>
              {organization.description}
            </p>
            <div>
              {organization.categories.map(categorie => (
                <Icon icon={categorie} />
              ))}
            </div>
          </div>
        </Content>
      </MediaContent>
      <MediaRight className="interestedContent">
        <Button color="isPrimary">
          Tenho interesse
        </Button>
      </MediaRight>
    </Media>
  </Container>
)

export default ClassifiedCard